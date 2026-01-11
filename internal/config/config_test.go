package config

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestLoad_MissingConfigEmitsWarning(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.toml")
	cfg, err := Load(path)
	if err == nil || !errors.Is(err, ErrConfigMissing) {
		t.Fatalf("expected missing config error, got %v", err)
	}
	defaultCfg := Default()
	if cfg.UI.Theme != defaultCfg.UI.Theme {
		t.Fatalf("expected default theme, got %q", cfg.UI.Theme)
	}
}

func TestLoad_InvalidConfigEmitsWarning(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.toml")
	if err := os.WriteFile(path, []byte("invalid toml = ="), 0o644); err != nil {
		t.Fatalf("failed to write invalid config: %v", err)
	}

	cfg, err := Load(path)
	if err == nil || !errors.Is(err, ErrConfigInvalid) {
		t.Fatalf("expected invalid config error, got %v", err)
	}
	if cfg.UI.Theme != Default().UI.Theme {
		t.Fatalf("expected default theme, got %q", cfg.UI.Theme)
	}
}
