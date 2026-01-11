package registry

import (
	"os"
	"path/filepath"
	"testing"
)

func TestListThemes_DeterministicAndFiltered(t *testing.T) {
	temp := t.TempDir()
	registryRoot := filepath.Join(temp, "registry")
	themesDir := filepath.Join(registryRoot, "themes")
	if err := os.MkdirAll(themesDir, 0o755); err != nil {
		t.Fatalf("failed to create themes dir: %v", err)
	}

	createTheme(t, themesDir, "beta", true)
	createTheme(t, themesDir, "alpha", true)
	createTheme(t, themesDir, "hidden", false)
	createFile(t, themesDir, "global-file", []byte("set bg black"))

	names, err := ListThemes(registryRoot)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []string{"alpha", "beta"}
	if len(names) != len(want) {
		t.Fatalf("unexpected theme count: got %v want %v", names, want)
	}
	for i := range want {
		if names[i] != want[i] {
			t.Fatalf("unexpected theme order: got %v want %v", names, want)
		}
	}
}

func TestListThemes_MissingThemesDir(t *testing.T) {
	temp := t.TempDir()
	registryRoot := filepath.Join(temp, "registry")

	names, err := ListThemes(registryRoot)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(names) != 0 {
		t.Fatalf("expected no themes, got %v", names)
	}
}

func TestListThemes_RepoSeed(t *testing.T) {
	root, err := findRepoRoot()
	if err != nil {
		t.Fatalf("failed to locate repo root: %v", err)
	}
	themesDir := filepath.Join(root, "registry")

	names, err := ListThemes(themesDir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	found := false
	for _, name := range names {
		if name == "default-dark" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("default-dark theme missing from registry list: %v", names)
	}
}

func createTheme(t *testing.T, themesDir, name string, withColors bool) {
	t.Helper()
	dir := filepath.Join(themesDir, name)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatalf("failed to create theme dir %s: %v", name, err)
	}
	if withColors {
		path := filepath.Join(dir, "colors")
		if err := os.WriteFile(path, []byte("set bg black\n"), 0o644); err != nil {
			t.Fatalf("failed to write colors for %s: %v", name, err)
		}
	}
}

func createFile(t *testing.T, dir, name string, contents []byte) {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, contents, 0o644); err != nil {
		t.Fatalf("failed to write file %s: %v", name, err)
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
