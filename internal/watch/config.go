package watch

import (
	"encoding/json"
	"os"
	"path/filepath"
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
	json.Unmarshal(data, &dirs)
	return dirs, nil
}

func SaveDirs(dirs []string) error {
	path := ConfigPath()
	os.MkdirAll(filepath.Dir(path), 0o755)
	data, _ := json.MarshalIndent(dirs, "", "  ")
	return os.WriteFile(path, data, 0o644)
}
