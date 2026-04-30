package tui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	s := "Shortcuts managed by Doppio\n\n"

	for i, shortcut := range m.shortcuts {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, shortcut, m.mode)

	}

	s += "\nPress q to quit.\n"

	return tea.NewView(s)
}
