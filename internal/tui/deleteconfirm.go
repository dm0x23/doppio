package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/dm0x23/doppio/internal/storage"
)

type DeleteConfirmModel struct {
	shortcutName string
}

func NewDeleteConfirmModel(name string) DeleteConfirmModel {
	return DeleteConfirmModel{shortcutName: name}
}

func (dc *DeleteConfirmModel) Update(msg tea.Msg) tea.Cmd {
	key, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return nil
	}
	switch key.String() {
	case "y":
		storage.Remove(dc.shortcutName)
		return func() tea.Msg { return shortcutDeletedMsg{} }
	case "n", "esc":
		return func() tea.Msg { return backToListMsg{} }
	}
	return nil
}

func (dc DeleteConfirmModel) View() string {
	return "Delete '" + dc.shortcutName + "'?\n\n(y)es  (n)o"
}
