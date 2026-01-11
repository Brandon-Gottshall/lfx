package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

const CurrentVersion = 1

type Config struct {
	ConfigVersion int             `toml:"config_version"`
	UI            UIConfig        `toml:"ui"`
	Extensions    map[string]bool `toml:"extensions"`
	Paths         PathsConfig     `toml:"paths"`
}

var (
	ErrConfigMissing = errors.New("config missing")
	ErrConfigInvalid = errors.New("config invalid")
)

type UIConfig struct {
	Theme string `toml:"theme"`
	Icons string `toml:"icons"`
	Font  string `toml:"font"`
}

type PathsConfig struct {
	Themes     string `toml:"themes"`
	Icons      string `toml:"icons"`
	Extensions string `toml:"extensions"`
}

func Default() Config {
	return Config{
		ConfigVersion: CurrentVersion,
		UI: UIConfig{
			Theme: "default-dark",
			Icons: "nerd-fonts",
			Font:  "monospace",
		},
		Extensions: map[string]bool{
			"hotkeys-hud": true,
		},
		Paths: PathsConfig{},
	}
}

func Load(configPath string) (Config, error) {
	cfg := Default()
	if configPath == "" {
		return cfg, nil
	}

	info, err := os.Stat(configPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return cfg, fmt.Errorf("%w: %s", ErrConfigMissing, configPath)
		}
		return cfg, err
	}
	if info.IsDir() {
		return cfg, fmt.Errorf("%w: config path is a directory", ErrConfigInvalid)
	}

	if _, err := toml.DecodeFile(configPath, &cfg); err != nil {
		return cfg, fmt.Errorf("%w: %w", ErrConfigInvalid, err)
	}

	if cfg.ConfigVersion == 0 {
		cfg.ConfigVersion = CurrentVersion
	}
	if cfg.ConfigVersion != CurrentVersion {
		return cfg, fmt.Errorf("%w: unsupported config_version %d", ErrConfigInvalid, cfg.ConfigVersion)
	}

	return cfg, nil
}

func DefaultPath(configDir string) string {
	if configDir == "" {
		return ""
	}
	return filepath.Join(configDir, "config.toml")
}
