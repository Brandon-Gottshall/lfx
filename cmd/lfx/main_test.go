package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestThemeListCLI(t *testing.T) {
	repoRoot, err := findRepoRoot()
	if err != nil {
		t.Fatalf("failed to locate repo root: %v", err)
	}

	cmd := exec.Command("go", "run", "./cmd/lfx", "theme", "list")
	cmd.Dir = repoRoot
	cmd.Env = append(os.Environ(), "GOCACHE="+filepath.Join(t.TempDir(), "go-build"))

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("lfx theme list failed: %v\n%s", err, output)
	}

	got := string(output)
	if !strings.Contains(got, "Themes") {
		t.Fatalf("expected Themes header, got: %s", got)
	}
	if !strings.Contains(got, "- default-dark") {
		t.Fatalf("expected default-dark theme to be listed, got: %s", got)
	}
}

func TestThemeSetCLI(t *testing.T) {
	repoRoot, err := findRepoRoot()
	if err != nil {
		t.Fatalf("failed to locate repo root: %v", err)
	}

	configDir := filepath.Join(t.TempDir(), "config")
	cmd := exec.Command("go", "run", "./cmd/lfx", "theme", "set", "default-dark")
	cmd.Dir = repoRoot
	cmd.Env = append(os.Environ(),
		"GOCACHE="+filepath.Join(t.TempDir(), "go-build"),
		"XDG_CONFIG_HOME="+configDir,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("lfx theme set failed: %v\n%s", err, output)
	}

	configRoot := filepath.Join(configDir, "lf")
	colorsPath := filepath.Join(configRoot, "colors")
	if _, err := os.Stat(colorsPath); err != nil {
		t.Fatalf("colors file missing: %v", err)
	}

	lfrcPath := filepath.Join(configRoot, "lfrc")
	data, err := os.ReadFile(lfrcPath)
	if err != nil {
		t.Fatalf("lfrc missing: %v", err)
	}
	if !strings.Contains(string(data), "# lfx:begin") {
		t.Fatalf("expected managed block, got: %s", string(data))
	}
}

func TestPluginListCLI(t *testing.T) {
	repoRoot, err := findRepoRoot()
	if err != nil {
		t.Fatalf("failed to locate repo root: %v", err)
	}

	cmd := exec.Command("go", "run", "./cmd/lfx", "plugin", "list")
	cmd.Dir = repoRoot
	cmd.Env = append(os.Environ(), "GOCACHE="+filepath.Join(t.TempDir(), "go-build"))

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("lfx plugin list failed: %v\n%s", err, output)
	}

	got := string(output)
	if !strings.Contains(got, "Plugins") {
		t.Fatalf("expected Plugins header, got: %s", got)
	}
	if !strings.Contains(got, "- hotkeys-hud") {
		t.Fatalf("expected hotkeys-hud plugin to be listed, got: %s", got)
	}
}

func TestPluginInstallCLI(t *testing.T) {
	repoRoot, err := findRepoRoot()
	if err != nil {
		t.Fatalf("failed to locate repo root: %v", err)
	}

	configDir := filepath.Join(t.TempDir(), "config")
	cmd := exec.Command("go", "run", "./cmd/lfx", "plugin", "install", "hotkeys-hud")
	cmd.Dir = repoRoot
	cmd.Env = append(os.Environ(),
		"GOCACHE="+filepath.Join(t.TempDir(), "go-build"),
		"XDG_CONFIG_HOME="+configDir,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("lfx plugin install failed: %v\n%s", err, output)
	}

	pluginPath := filepath.Join(configDir, "lf", "plugins", "hotkeys-hud.lfrc")
	if _, err := os.Stat(pluginPath); err != nil {
		t.Fatalf("plugin file missing: %v", err)
	}

	lfrcPath := filepath.Join(configDir, "lf", "lfrc")
	data, err := os.ReadFile(lfrcPath)
	if err != nil {
		t.Fatalf("lfrc missing: %v", err)
	}
	if !strings.Contains(string(data), "source "+pluginPath) {
		t.Fatalf("expected plugin source line, got: %s", string(data))
	}
}

func TestPluginUninstallCLI(t *testing.T) {
	repoRoot, err := findRepoRoot()
	if err != nil {
		t.Fatalf("failed to locate repo root: %v", err)
	}

	configDir := filepath.Join(t.TempDir(), "config")
	installCmd := exec.Command("go", "run", "./cmd/lfx", "plugin", "install", "hotkeys-hud")
	installCmd.Dir = repoRoot
	installCmd.Env = append(os.Environ(),
		"GOCACHE="+filepath.Join(t.TempDir(), "go-build"),
		"XDG_CONFIG_HOME="+configDir,
	)

	if output, err := installCmd.CombinedOutput(); err != nil {
		t.Fatalf("lfx plugin install failed: %v\n%s", err, output)
	}

	uninstallCmd := exec.Command("go", "run", "./cmd/lfx", "plugin", "uninstall", "hotkeys-hud")
	uninstallCmd.Dir = repoRoot
	uninstallCmd.Env = append(os.Environ(),
		"GOCACHE="+filepath.Join(t.TempDir(), "go-build"),
		"XDG_CONFIG_HOME="+configDir,
	)

	output, err := uninstallCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("lfx plugin uninstall failed: %v\n%s", err, output)
	}

	pluginPath := filepath.Join(configDir, "lf", "plugins", "hotkeys-hud.lfrc")
	if _, err := os.Stat(pluginPath); err == nil {
		t.Fatalf("expected plugin file to be removed")
	}

	lfrcPath := filepath.Join(configDir, "lf", "lfrc")
	data, err := os.ReadFile(lfrcPath)
	if err != nil {
		t.Fatalf("lfrc missing: %v", err)
	}
	if strings.Contains(string(data), "source "+pluginPath) {
		t.Fatalf("expected plugin source line removed, got: %s", string(data))
	}
}

func findRepoRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", os.ErrNotExist
		}
		dir = parent
	}
}
