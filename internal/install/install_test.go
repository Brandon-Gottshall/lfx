package install

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestApplyTheme_CopiesColorsAndUpdatesLfrc(t *testing.T) {
	temp := t.TempDir()
	registry := filepath.Join(temp, "registry")
	themeName := "default-dark"
	colors := "set bg black\n"
	createTheme(t, registry, themeName, colors)

	lfDir := filepath.Join(temp, "lf")
	if err := ApplyTheme(registry, themeName, lfDir); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(lfDir, "colors"))
	if err != nil {
		t.Fatalf("failed to read colors file: %v", err)
	}
	if string(data) != colors {
		t.Fatalf("colors mismatch: got %q want %q", string(data), colors)
	}

	lfrc, err := os.ReadFile(filepath.Join(lfDir, "lfrc"))
	if err != nil {
		t.Fatalf("failed to read lfrc: %v", err)
	}
	expected, err := buildManagedBlock(lfDir)
	if err != nil {
		t.Fatalf("failed to build managed block: %v", err)
	}
	if string(lfrc) != expected {
		t.Fatalf("unexpected lfrc\n%s", string(lfrc))
	}
}

func TestApplyTheme_Idempotent(t *testing.T) {
	temp := t.TempDir()
	registry := filepath.Join(temp, "registry")
	themeName := "cobalt"
	colors := "set fg cyan\n"
	createTheme(t, registry, themeName, colors)

	lfDir := filepath.Join(temp, "lf")
	if err := ApplyTheme(registry, themeName, lfDir); err != nil {
		t.Fatalf("first apply error: %v", err)
	}
	first, err := os.ReadFile(filepath.Join(lfDir, "lfrc"))
	if err != nil {
		t.Fatalf("failed to read lfrc: %v", err)
	}

	if err := ApplyTheme(registry, themeName, lfDir); err != nil {
		t.Fatalf("second apply error: %v", err)
	}
	second, err := os.ReadFile(filepath.Join(lfDir, "lfrc"))
	if err != nil {
		t.Fatalf("failed to read lfrc second time: %v", err)
	}

	if string(first) != string(second) {
		t.Fatalf("lfrc changed between runs\nfirst:\n%s\nsecond:\n%s", first, second)
	}
}

func TestApplyTheme_ReplacesExistingBlock(t *testing.T) {
	temp := t.TempDir()
	registry := filepath.Join(temp, "registry")
	themeName := "solar"
	colors := "set bg solarized\n"
	createTheme(t, registry, themeName, colors)

	lfDir := filepath.Join(temp, "lf")
	if err := os.MkdirAll(lfDir, 0o755); err != nil {
		t.Fatalf("failed to prepare lf dir: %v", err)
	}
	lfrcPath := filepath.Join(lfDir, "lfrc")

	oldBlock := fmt.Sprintf("%s\n%s\nset colors /tmp/old/colors\n%s\n", blockBegin, blockManaged, blockEnd)
	if err := os.WriteFile(lfrcPath, []byte("intro\n"+oldBlock+"outro\n"), 0o644); err != nil {
		t.Fatalf("failed to write old lfrc: %v", err)
	}

	if err := ApplyTheme(registry, themeName, lfDir); err != nil {
		t.Fatalf("apply returned error: %v", err)
	}

	data, err := os.ReadFile(lfrcPath)
	if err != nil {
		t.Fatalf("failed to read lfrc: %v", err)
	}
	if strings.Count(string(data), blockBegin) != 1 {
		t.Fatalf("expected single block, got %s", string(data))
	}
	if !strings.Contains(string(data), "intro") || !strings.Contains(string(data), "outro") {
		t.Fatalf("context was lost: %s", string(data))
	}
}

func TestApplyTheme_MissingTheme(t *testing.T) {
	temp := t.TempDir()
	registry := filepath.Join(temp, "registry")
	lfDir := filepath.Join(temp, "lf")

	if err := ApplyTheme(registry, "missing", lfDir); err == nil {
		t.Fatal("expected error when theme missing")
	}
}

func TestApplyPlugin_CopiesSnippetAndUpdatesLfrc(t *testing.T) {
	temp := t.TempDir()
	registry := filepath.Join(temp, "registry")
	pluginName := "hotkeys-hud"
	snippet := "# plugin snippet\n"
	createPlugin(t, registry, pluginName, snippet)

	lfDir := filepath.Join(temp, "lf")
	if err := ApplyPlugin(registry, pluginName, lfDir); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(lfDir, "plugins", pluginName+".lfrc"))
	if err != nil {
		t.Fatalf("failed to read plugin file: %v", err)
	}
	if string(data) != snippet {
		t.Fatalf("plugin mismatch: got %q want %q", string(data), snippet)
	}

	lfrc, err := os.ReadFile(filepath.Join(lfDir, "lfrc"))
	if err != nil {
		t.Fatalf("failed to read lfrc: %v", err)
	}
	expected, err := buildManagedBlock(lfDir)
	if err != nil {
		t.Fatalf("failed to build managed block: %v", err)
	}
	if string(lfrc) != expected {
		t.Fatalf("unexpected lfrc\n%s", string(lfrc))
	}
	if !strings.Contains(string(lfrc), "source "+filepath.Join(lfDir, "plugins", pluginName+".lfrc")) {
		t.Fatalf("expected plugin source line, got %s", string(lfrc))
	}
}

func TestApplyPlugin_Idempotent(t *testing.T) {
	temp := t.TempDir()
	registry := filepath.Join(temp, "registry")
	pluginName := "hotkeys-hud"
	snippet := "# plugin snippet\n"
	createPlugin(t, registry, pluginName, snippet)

	lfDir := filepath.Join(temp, "lf")
	if err := ApplyPlugin(registry, pluginName, lfDir); err != nil {
		t.Fatalf("first install error: %v", err)
	}
	first, err := os.ReadFile(filepath.Join(lfDir, "lfrc"))
	if err != nil {
		t.Fatalf("failed to read lfrc: %v", err)
	}

	if err := ApplyPlugin(registry, pluginName, lfDir); err != nil {
		t.Fatalf("second install error: %v", err)
	}
	second, err := os.ReadFile(filepath.Join(lfDir, "lfrc"))
	if err != nil {
		t.Fatalf("failed to read lfrc second time: %v", err)
	}

	if string(first) != string(second) {
		t.Fatalf("lfrc changed between runs\nfirst:\n%s\nsecond:\n%s", first, second)
	}
}

func TestApplyPlugin_MissingPlugin(t *testing.T) {
	temp := t.TempDir()
	registry := filepath.Join(temp, "registry")
	lfDir := filepath.Join(temp, "lf")

	if err := ApplyPlugin(registry, "missing", lfDir); err == nil {
		t.Fatal("expected error when plugin missing")
	}
}

func TestRemovePlugin_RemovesSnippetAndUpdatesLfrc(t *testing.T) {
	temp := t.TempDir()
	registry := filepath.Join(temp, "registry")
	pluginName := "hotkeys-hud"
	snippet := "# plugin snippet\n"
	createPlugin(t, registry, pluginName, snippet)

	lfDir := filepath.Join(temp, "lf")
	if err := ApplyPlugin(registry, pluginName, lfDir); err != nil {
		t.Fatalf("install returned error: %v", err)
	}

	if err := RemovePlugin(lfDir, pluginName); err != nil {
		t.Fatalf("remove returned error: %v", err)
	}

	if _, err := os.Stat(filepath.Join(lfDir, "plugins", pluginName+".lfrc")); err == nil {
		t.Fatal("expected plugin file to be removed")
	}

	lfrc, err := os.ReadFile(filepath.Join(lfDir, "lfrc"))
	if err != nil {
		t.Fatalf("failed to read lfrc: %v", err)
	}
	if strings.Contains(string(lfrc), "source "+filepath.Join(lfDir, "plugins", pluginName+".lfrc")) {
		t.Fatalf("expected plugin source line to be removed, got %s", string(lfrc))
	}
}

func TestRemovePlugin_Idempotent(t *testing.T) {
	temp := t.TempDir()
	lfDir := filepath.Join(temp, "lf")

	if err := RemovePlugin(lfDir, "missing"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRemovePlugin_PreservesContextAndRewritesBlock(t *testing.T) {
	temp := t.TempDir()
	lfDir := filepath.Join(temp, "lf")
	pluginsDir := filepath.Join(lfDir, "plugins")
	if err := os.MkdirAll(pluginsDir, 0o755); err != nil {
		t.Fatalf("failed to create plugins dir: %v", err)
	}

	pluginName := "hotkeys-hud"
	pluginPath := filepath.Join(pluginsDir, pluginName+".lfrc")
	if err := os.WriteFile(pluginPath, []byte("# plugin snippet\n"), 0o644); err != nil {
		t.Fatalf("failed to write plugin file: %v", err)
	}

	colorsPath := filepath.Join(lfDir, "colors")
	if err := os.WriteFile(colorsPath, []byte("set bg black\n"), 0o644); err != nil {
		t.Fatalf("failed to write colors file: %v", err)
	}

	block, err := buildManagedBlock(lfDir)
	if err != nil {
		t.Fatalf("failed to build managed block: %v", err)
	}
	lfrcPath := filepath.Join(lfDir, "lfrc")
	if err := os.WriteFile(lfrcPath, []byte("intro\n"+block+"outro\n"), 0o644); err != nil {
		t.Fatalf("failed to write lfrc: %v", err)
	}

	if err := RemovePlugin(lfDir, pluginName); err != nil {
		t.Fatalf("remove returned error: %v", err)
	}

	data, err := os.ReadFile(lfrcPath)
	if err != nil {
		t.Fatalf("failed to read lfrc: %v", err)
	}
	if strings.Count(string(data), blockBegin) != 1 {
		t.Fatalf("expected single block, got %s", string(data))
	}
	if !strings.Contains(string(data), "intro") || !strings.Contains(string(data), "outro") {
		t.Fatalf("context was lost: %s", string(data))
	}
	if strings.Contains(string(data), "source "+pluginPath) {
		t.Fatalf("expected plugin source line removed, got %s", string(data))
	}
	if !strings.Contains(string(data), "# colors file at "+colorsPath) {
		t.Fatalf("expected colors line to remain, got %s", string(data))
	}
}

func createTheme(t *testing.T, registry, name, contents string) {
	t.Helper()
	path := filepath.Join(registry, "themes", name)
	if err := os.MkdirAll(path, 0o755); err != nil {
		t.Fatalf("failed to create theme dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(path, "colors"), []byte(contents), 0o644); err != nil {
		t.Fatalf("failed to write colors file: %v", err)
	}
}

func createPlugin(t *testing.T, registry, name, contents string) {
	t.Helper()
	path := filepath.Join(registry, "plugins")
	if err := os.MkdirAll(path, 0o755); err != nil {
		t.Fatalf("failed to create plugins dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(path, name+".lfrc"), []byte(contents), 0o644); err != nil {
		t.Fatalf("failed to write plugin file: %v", err)
	}
}
