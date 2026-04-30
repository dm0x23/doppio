package tui

import (
	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.shortcuts)-1 {
				m.cursor++
			}

		case "a":
			m.mode = "add"

		case "r":
			m.mode = "remove"

		case "s":
			m.mode = "sync"
		}
	}

	return m, nil
}
