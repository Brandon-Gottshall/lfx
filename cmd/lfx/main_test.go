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
