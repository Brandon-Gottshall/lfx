package config

import (
	"errors"
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
			return cfg, nil
		}
		return cfg, err
	}
	if info.IsDir() {
		return cfg, errors.New("config path is a directory")
	}

	if _, err := toml.DecodeFile(configPath, &cfg); err != nil {
		return cfg, err
	}

	if cfg.ConfigVersion == 0 {
		cfg.ConfigVersion = CurrentVersion
	}
	if cfg.ConfigVersion != CurrentVersion {
		return cfg, errors.New("unsupported config_version")
	}

	return cfg, nil
}

func DefaultPath(configDir string) string {
	if configDir == "" {
		return ""
	}
	return filepath.Join(configDir, "config.toml")
}
