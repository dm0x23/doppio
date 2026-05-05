package watch

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

func ConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "doppio", "watch.json")
}

func LoadDirs() ([]string, error) {
	data, err := os.ReadFile(ConfigPath())
	if err != nil {
		return []string{}, nil
	}
	var dirs []string
	if err := json.Unmarshal(data, &dirs); err != nil {
		return []string{}, nil
	}
	var expanded []string
	for _, d := range dirs {
		expanded = append(expanded, expandHome(d))
	}
	return expanded, nil
}

func SaveDirs(dirs []string) error {
	var expanded []string
	for _, d := range dirs {
		expanded = append(expanded, expandHome(d))
	}
	path := ConfigPath()
	os.MkdirAll(filepath.Dir(path), 0o755)
	data, _ := json.MarshalIndent(expanded, "", "  ")
	return os.WriteFile(path, data, 0o644)
}

func expandHome(path string) string {
	if strings.HasPrefix(path, "~/") || path == "~" {
		home, err := os.UserHomeDir()
		if err == nil {
			return filepath.Join(home, path[1:])
		}
	}
	return path
}
