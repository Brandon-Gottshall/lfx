package install

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	blockBegin   = "# lfx:begin"
	blockManaged = "# lfx:managed"
	blockEnd     = "# lfx:end"
)

func ApplyTheme(registryRoot, themeName, lfConfigDir string) error {
	if registryRoot == "" {
		return errors.New("registry root not provided")
	}
	colorsSrc := filepath.Join(registryRoot, "themes", themeName, "colors")
	if _, err := os.Stat(colorsSrc); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("theme not found: %s", themeName)
		}
		return fmt.Errorf("failed to access theme %s: %w", themeName, err)
	}

	if lfConfigDir == "" {
		return errors.New("lf config dir not provided")
	}
	if err := os.MkdirAll(lfConfigDir, 0o755); err != nil {
		return fmt.Errorf("failed to create lf config dir: %w", err)
	}

	colorsDst := filepath.Join(lfConfigDir, "colors")
	if err := copyFile(colorsSrc, colorsDst); err != nil {
		return fmt.Errorf("failed to copy colors file: %w", err)
	}

	block, err := buildManagedBlock(lfConfigDir)
	if err != nil {
		return fmt.Errorf("failed to build managed block: %w", err)
	}
	if err := writeManagedLfrc(filepath.Join(lfConfigDir, "lfrc"), block); err != nil {
		return fmt.Errorf("failed to update lfrc: %w", err)
	}

	return nil
}

func ApplyPlugin(registryRoot, pluginName, lfConfigDir string) error {
	if registryRoot == "" {
		return errors.New("registry root not provided")
	}
	pluginSrc := filepath.Join(registryRoot, "plugins", pluginName+".lfrc")
	if _, err := os.Stat(pluginSrc); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("plugin not found: %s", pluginName)
		}
		return fmt.Errorf("failed to access plugin %s: %w", pluginName, err)
	}

	if lfConfigDir == "" {
		return errors.New("lf config dir not provided")
	}
	if err := os.MkdirAll(lfConfigDir, 0o755); err != nil {
		return fmt.Errorf("failed to create lf config dir: %w", err)
	}

	pluginsDir := filepath.Join(lfConfigDir, "plugins")
	if err := os.MkdirAll(pluginsDir, 0o755); err != nil {
		return fmt.Errorf("failed to create plugins dir: %w", err)
	}
	pluginDst := filepath.Join(pluginsDir, pluginName+".lfrc")
	if err := copyFile(pluginSrc, pluginDst); err != nil {
		return fmt.Errorf("failed to copy plugin file: %w", err)
	}

	block, err := buildManagedBlock(lfConfigDir)
	if err != nil {
		return fmt.Errorf("failed to build managed block: %w", err)
	}
	if err := writeManagedLfrc(filepath.Join(lfConfigDir, "lfrc"), block); err != nil {
		return fmt.Errorf("failed to update lfrc: %w", err)
	}

	return nil
}

func RemovePlugin(lfConfigDir, pluginName string) error {
	if lfConfigDir == "" {
		return errors.New("lf config dir not provided")
	}
	if pluginName == "" {
		return errors.New("plugin name not provided")
	}

	info, err := os.Stat(lfConfigDir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return fmt.Errorf("failed to access lf config dir: %w", err)
	}
	if !info.IsDir() {
		return fmt.Errorf("lf config dir is not a directory: %s", lfConfigDir)
	}

	pluginPath := filepath.Join(lfConfigDir, "plugins", pluginName+".lfrc")
	if err := os.Remove(pluginPath); err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("failed to remove plugin file: %w", err)
	}

	block, err := buildManagedBlock(lfConfigDir)
	if err != nil {
		return fmt.Errorf("failed to build managed block: %w", err)
	}
	if err := writeManagedLfrc(filepath.Join(lfConfigDir, "lfrc"), block); err != nil {
		return fmt.Errorf("failed to update lfrc: %w", err)
	}

	return nil
}

func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0o644)
}

func buildManagedBlock(lfConfigDir string) (string, error) {
	lines := []string{blockBegin, blockManaged}

	colorsPath := filepath.Join(lfConfigDir, "colors")
	if _, err := os.Stat(colorsPath); err == nil {
		lines = append(lines, fmt.Sprintf("# colors file at %s", colorsPath))
	} else if !errors.Is(err, os.ErrNotExist) {
		return "", err
	}

	pluginFiles, err := listPluginFiles(filepath.Join(lfConfigDir, "plugins"))
	if err != nil {
		return "", err
	}
	if len(pluginFiles) > 0 {
		lines = append(lines, "# plugins")
		for _, plugin := range pluginFiles {
			lines = append(lines, fmt.Sprintf("source %s", plugin))
		}
	}

	lines = append(lines, blockEnd, "")
	return strings.Join(lines, "\n"), nil
}

func writeManagedLfrc(path, block string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}

	updated := setManagedBlock(string(data), block)
	return os.WriteFile(path, []byte(updated), 0o644)
}

func listPluginFiles(pluginDir string) ([]string, error) {
	entries, err := os.ReadDir(pluginDir)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return []string{}, nil
		}
		return nil, err
	}

	var plugins []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if strings.HasPrefix(name, ".") {
			continue
		}
		if filepath.Ext(name) != ".lfrc" {
			continue
		}
		plugins = append(plugins, filepath.Join(pluginDir, name))
	}

	sort.Strings(plugins)
	return plugins, nil
}

func setManagedBlock(existing, block string) string {
	cleaned := removeManagedBlock(existing)
	trimmed := strings.TrimRight(cleaned, "\n")
	if trimmed != "" {
		trimmed += "\n\n"
	}
	result := trimmed + block
	if !strings.HasSuffix(result, "\n") {
		result += "\n"
	}
	return result
}

func removeManagedBlock(content string) string {
	lines := strings.Split(content, "\n")
	var builder strings.Builder
	inBlock := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == blockBegin {
			inBlock = true
			continue
		}
		if inBlock {
			if trimmed == blockEnd {
				inBlock = false
			}
			continue
		}

		if builder.Len() > 0 {
			builder.WriteByte('\n')
		}
		builder.WriteString(line)
	}

	return builder.String()
}
