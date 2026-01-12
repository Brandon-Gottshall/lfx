package install

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
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

	block := buildLfrcBlock(colorsDst)
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

func buildLfrcBlock(colorsPath string) string {
	return strings.Join([]string{
		blockBegin,
		blockManaged,
		fmt.Sprintf("# colors file at %s", colorsPath),
		blockEnd,
		"",
	}, "\n")
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
