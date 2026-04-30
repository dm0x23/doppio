package tui

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type AddFormModel struct {
	nameInput    textinput.Model
	commandInput textinput.Model
	focusIndex   int
}

func NewAddFormModel() AddFormModel {
	name := textinput.New()
	name.Placeholder = "alias name"
	name.Focus()
	cmd := textinput.New()
	cmd.Placeholder = "command"
	return AddFormModel{
		nameInput:    name,
		commandInput: cmd,
		focusIndex:   0,
	}
}

func (af *AddFormModel) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "esc":
			return func() tea.Msg { return backToListMsg{} }
		case "tab":
			af.focusIndex = (af.focusIndex + 1) % 2
			af.focus()
		case "shift+tab":
			af.focusIndex = (af.focusIndex - 1 + 2) % 2
			af.focus()
		case "enter":
			name := af.nameInput.Value()
			cmd := af.commandInput.Value()
			if name != "" && cmd != "" {
				return func() tea.Msg { return newShortcutMsg{name, cmd} }
			}
		}
	}
	if af.focusIndex == 0 {
		var cmd tea.Cmd
		af.nameInput, cmd = af.nameInput.Update(msg)
		return cmd
	}
	var cmd tea.Cmd
	af.commandInput, cmd = af.commandInput.Update(msg)
	return cmd
}

func (af *AddFormModel) focus() {
	if af.focusIndex == 0 {
		af.nameInput.Focus()
		af.commandInput.Blur()
	} else {
		af.commandInput.Focus()
		af.nameInput.Blur()
	}
}

func (af AddFormModel) View() string {
	return "Add shortcut\n\n" +
		"Name:\n" + InputStyle.Render(af.nameInput.View()) + "\n\n" +
		"Command:\n" + InputStyle.Render(af.commandInput.View()) + "\n\n" +
		"Enter · Save   Esc · Back"
}
