package tui

import (
	"fmt"
	"os"
	"os/exec"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/dm0x23/doppio/internal/watch"
)

type WatchModel struct {
	dirs   []string
	cursor int
	input  textinput.Model
	adding bool
}

func NewWatchModel() WatchModel {
	dirs, _ := watch.LoadDirs()

	input := textinput.New()
	input.Placeholder = "Directory path"

	return WatchModel{
		dirs:   dirs,
		cursor: 0,
		input:  input,
		adding: false,
	}
}

func (wm *WatchModel) Update(msg tea.Msg) tea.Cmd {
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

					for _, d := range wm.dirs {
						if d == newDir {
							wm.adding = false
							wm.input.SetValue("")
							return nil
						}
					}
					wm.dirs = append(wm.dirs, newDir)
					if err := watch.SaveDirs(wm.dirs); err != nil {
						fmt.Fprintf(os.Stderr, "ERROR saving watch dirs: %v\n", err)
					}
				}
				restartWatchDaemon()
				wm.adding = false
				wm.input.SetValue("")
				return nil
			}

			var cmd tea.Cmd
			wm.input, cmd = wm.input.Update(msg)
			return cmd

		default:

			var cmd tea.Cmd
			wm.input, cmd = wm.input.Update(msg)
			return cmd
		}
	}

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
			if err := watch.SaveDirs(wm.dirs); err != nil {
				fmt.Fprintf(os.Stderr, "ERROR saving watch dirs: %v\n", err)
			}
		}
		restartWatchDaemon()
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

func restartWatchDaemon() error {
	if err := exec.Command("systemctl", "--user", "is-active", "--quiet", "doppio-watch").Run(); err != nil {
		return exec.Command("systemctl", "--user", "start", "doppio-watch").Run()
	}
	return exec.Command("systemctl", "--user", "restart", "doppio-watch").Run()
}
