package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/dm0x23/doppio/internal/storage"
)

type Model struct {
	screenStack []string
	cursor      int
	shortcuts   []storage.Shortcut
	width       int
	height      int

	addForm        AddFormModel
	deleteConfirm  DeleteConfirmModel
	watchModel     WatchModel
	bootstrapModel BootstrapModel

	statusMsg string
}

func NewModel() Model {
	shortcuts, _ := storage.Load()
	return Model{
		screenStack:    []string{"list"},
		shortcuts:      shortcuts,
		addForm:        NewAddFormModel(),
		deleteConfirm:  DeleteConfirmModel{},
		watchModel:     NewWatchModel(),
		bootstrapModel: NewBootstrapModel(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) pushScreen(s string) { m.screenStack = append(m.screenStack, s) }
func (m *Model) popScreen() string {
	if len(m.screenStack) == 0 {
		return ""
	}
	top := m.screenStack[len(m.screenStack)-1]
	m.screenStack = m.screenStack[:len(m.screenStack)-1]
	return top
}
func (m Model) currentScreen() string { return m.screenStack[len(m.screenStack)-1] }
