package registry

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func ListThemes(registryRoot string) ([]string, error) {
	themesDir := filepath.Join(registryRoot, "themes")
	entries, err := os.ReadDir(themesDir)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return []string{}, nil
		}
		return nil, err
	}

	var themes []string
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		name := entry.Name()
		if strings.HasPrefix(name, ".") {
			continue
		}
		colorsPath := filepath.Join(themesDir, name, "colors")
		if _, err := os.Stat(colorsPath); err != nil {
			continue
		}
		themes = append(themes, name)
	}

	sort.Strings(themes)
	return themes, nil
}
