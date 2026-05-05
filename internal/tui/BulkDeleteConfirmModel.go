package tui

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/dm0x23/doppio/internal/storage"
)

type BulkDeleteConfirmModel struct {
	names []string
}

func NewBulkDeleteConfirmModel(names []string) BulkDeleteConfirmModel {
	return BulkDeleteConfirmModel{names: names}
}

func (bd *BulkDeleteConfirmModel) Update(msg tea.Msg) tea.Cmd {
	key, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return nil
	}
	switch key.String() {
	case "y":
		for _, name := range bd.names {
			storage.Remove(name)
		}
		return func() tea.Msg { return shortcutDeletedMsg{} }
	case "n", "esc":
		return func() tea.Msg { return backToListMsg{} }
	}
	return nil
}

func (bd BulkDeleteConfirmModel) View() string {
	return "Delete " + strings.Join(bd.names, ", ") + "?\n\n(y)es  (n)o"
}
