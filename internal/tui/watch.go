package tui

import (
	"encoding/json"
	"os"
	"path/filepath"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type WatchModel struct {
	dirs   []string
	cursor int
	input  textinput.Model
	adding bool
}

func NewWatchModel() WatchModel {
	// Load saved watched dirs
	dirs, _ := loadWatchedDirs()

	input := textinput.New()
	input.Placeholder = "Directory path"
	input.Focus()

	return WatchModel{
		dirs:   dirs,
		cursor: 0,
		input:  input,
		adding: false,
	}
}

func (wm *WatchModel) Update(msg tea.Msg) tea.Cmd {
	// If we're in add mode, pass to textinput
	if wm.adding {
		switch msg := msg.(type) {
		case tea.KeyPressMsg:
			switch msg.String() {
			case "esc":
				wm.adding = false
				wm.input.Blur()
				wm.input.SetValue("")
				return nil
			case "enter":
				newDir := wm.input.Value()
				if newDir != "" {
					// Add directory only if it's not already watched
					for _, d := range wm.dirs {
						if d == newDir {
							wm.adding = false
							wm.input.SetValue("")
							return nil
						}
					}
					wm.dirs = append(wm.dirs, newDir)
					saveWatchedDirs(wm.dirs)
				}
				wm.adding = false
				wm.input.SetValue("")
				return nil
			}
		default:
			var cmd tea.Cmd
			wm.input, cmd = wm.input.Update(msg)
			return cmd
		}
		return nil
	}

	// Normal watch list mode
	key, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return nil
	}
	switch key.String() {
	case "j", "down":
		if wm.cursor < len(wm.dirs)-1 {
			wm.cursor++
		}
	case "k", "up":
		if wm.cursor > 0 {
			wm.cursor--
		}
	case "a":
		wm.adding = true
		wm.input.Focus()
		return nil
	case "d":
		if len(wm.dirs) > 0 {
			wm.dirs = append(wm.dirs[:wm.cursor], wm.dirs[wm.cursor+1:]...)
			if wm.cursor >= len(wm.dirs) && len(wm.dirs) > 0 {
				wm.cursor = len(wm.dirs) - 1
			}
			saveWatchedDirs(wm.dirs)
		}
	}
	return nil
}

func (wm WatchModel) View() string {
	if wm.adding {
		return "👀 Add directory to watch\n\n" +
			InputStyle.Render(wm.input.View()) + "\n\n" +
			"Enter · Save   Esc · Cancel"
	}

	s := "👀 Watched Directories\n\n"
	if len(wm.dirs) == 0 {
		s += "No directories watched.\n"
	}
	for i, d := range wm.dirs {
		cursor := "  "
		if wm.cursor == i {
			cursor = CursorMarker.Render(">") + " "
		}
		s += cursor + d + "\n"
	}
	s += "\n(a)dd  (d)elete  Esc · Back"
	return s
}

// ---- Persistence helpers (inline for simplicity) ----
func watchConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "doppio", "watch.json")
}

func loadWatchedDirs() ([]string, error) {
	data, err := os.ReadFile(watchConfigPath())
	if err != nil {
		return []string{}, nil
	}
	var dirs []string
	json.Unmarshal(data, &dirs)
	return dirs, nil
}

func saveWatchedDirs(dirs []string) error {
	path := watchConfigPath()
	os.MkdirAll(filepath.Dir(path), 0o755)
	data, _ := json.MarshalIndent(dirs, "", "  ")
	return os.WriteFile(path, data, 0o644)
}
