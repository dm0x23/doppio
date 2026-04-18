package storage

import (
	"encoding/json"
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
		os.IsNotExist(err){
			return []Shortcut{}, nil
		}
		return nil, err
	}
	var shortcuts []Shortcut
	json.Unmarshal(jsonContent, &shortcuts); err != nil{
		return nil, fmt.Errorf("failed to parse shortcuts.json %w", err)
	}

	return shortcuts, nil

}
