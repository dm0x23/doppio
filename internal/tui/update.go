package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/dm0x23/doppio/internal/storage"
	"github.com/dm0x23/doppio/internal/sync"
)

type (
	newShortcutMsg     struct{ name, command string }
	backToListMsg      struct{}
	shortcutDeletedMsg struct{}
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyPressMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "esc" && m.currentScreen() != "list" {
			m.popScreen()
			m.statusMsg = ""
			return m, nil
		}

	case newShortcutMsg:
		storage.Add(msg.name, msg.command, []string{"zsh", "bash"})
		sync.Run()
		m.shortcuts, _ = storage.List()
		m.popScreen()
		m.statusMsg = "Added " + msg.name
		return m, nil

	case backToListMsg:
		m.popScreen()
		return m, nil

	case shortcutDeletedMsg:
		sync.Run()
		m.shortcuts, _ = storage.List()
		m.popScreen()
		m.statusMsg = "Shortcut deleted"
		return m, nil
	}

	switch m.currentScreen() {
	case "list":
		return m.handleListKeys(msg)
	case "add":
		cmd := m.addForm.Update(msg)
		return m, cmd
	case "delete-confirm":
		cmd := m.deleteConfirm.Update(msg)
		return m, cmd
	case "watch":
		cmd := m.watchModel.Update(msg)
		return m, cmd
	case "bootstrap":
		cmd := m.bootstrapModel.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m Model) handleListKeys(msg tea.Msg) (tea.Model, tea.Cmd) {
	key, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return m, nil
	}
	switch key.String() {
	case "j", "down":
		if m.cursor < len(m.shortcuts)-1 {
			m.cursor++
		}
	case "k", "up":
		if m.cursor > 0 {
			m.cursor--
		}
	case "a":
		m.pushScreen("add")
		m.addForm = NewAddFormModel()
		return m, nil
	case "d":
		if len(m.shortcuts) == 0 {
			return m, nil
		}
		m.pushScreen("delete-confirm")
		m.deleteConfirm = NewDeleteConfirmModel(m.shortcuts[m.cursor].Name)
		return m, nil
	case "w":
		m.pushScreen("watch")
		return m, nil
	case "b":
		m.pushScreen("bootstrap")
		return m, nil
	case "s":
		sync.Run()
		m.shortcuts, _ = storage.List()
		m.statusMsg = "Synced"
		return m, nil
	}
	return m, nil
}
