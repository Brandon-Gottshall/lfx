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
	expected := buildLfrcBlock(filepath.Join(lfDir, "colors"))
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
