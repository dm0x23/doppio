package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/dm0x23/doppio/internal/storage"
)

type Model struct {
	shortcuts []storage.Shortcut
	cursor    int
	mode      string
}

func initModel() Model {
	Sshortcuts, _ := storage.Load()
	return Model{
		shortcuts: Sshortcuts,
		mode:      "list",
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
