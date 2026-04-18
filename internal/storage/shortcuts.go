package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Shortcut struct {
	Name        string    `json:"name"`
	Command     string    `json:"command"`
	Description string    `json:"description,omitempty"`
	Shells      []string  `json:"shells"`
	Created     time.Time `json:"created"`
}

func configPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, "doppio", "shortcuts.json"), nil
}

func Load() ([]Shortcut, error) {
	path, err := configPath()
	if err != nil {
		return nil, err
	}
	jsonContent, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Shortcut{}, nil
		}
		return nil, err
	}
	var shortcuts []Shortcut
	if json.Unmarshal(jsonContent, &shortcuts); err != nil {
		return nil, fmt.Errorf("failed to parse shortcuts.json %w", err)
	}

	return shortcuts, nil

}

func Save(shortcuts []Shortcut) error {
	path, err := configPath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o775); err != nil {
		return err
	}

	data, err := json.MarshalIndent(shortcuts, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
