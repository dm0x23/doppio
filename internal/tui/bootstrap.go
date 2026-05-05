package tui

import (
	"os/exec"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type ToolStatus struct {
	Name      string
	Installed bool
}

type BootstrapModel struct {
	tools   []ToolStatus
	cursor  int
	msg     string
	bootRan bool
}

func NewBootstrapModel() BootstrapModel {
	tools := []string{"zoxide", "fzf", "bat", "ripgrep", "eza"}
	statuses := make([]ToolStatus, len(tools))
	for i, t := range tools {
		_, err := exec.LookPath(t)
		// bat might be batcat on Ubuntu
		if t == "bat" && err != nil {
			_, err = exec.LookPath("batcat")
		}
		statuses[i] = ToolStatus{Name: t, Installed: err == nil}
	}
	return BootstrapModel{
		tools:  statuses,
		cursor: 0,
	}
}

func (bm *BootstrapModel) Update(msg tea.Msg) tea.Cmd {
	key, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return nil
	}
	switch key.String() {
	case "j", "down":
		if bm.cursor < len(bm.tools)-1 {
			bm.cursor++
		}
	case "k", "up":
		if bm.cursor > 0 {
			bm.cursor--
		}
	case "enter":
		if !bm.bootRan {
			bm.bootRan = true
			bm.msg = "Run 'dop bootstrap' in your terminal to install missing tools."
		}
	}
	return nil
}

func (bm BootstrapModel) View() string {
	s := "🛠️ Developer Tools Status\n\n"
	for i, tool := range bm.tools {
		cursor := "  "
		if bm.cursor == i {
			cursor = CursorMarker.Render(">") + " "
		}
		status := "✖ needs install"
		if tool.Installed {
			status = "✔ installed"
		}
		s += cursor + tool.Name + "  " + styleStatus(status, tool.Installed) + "\n"
	}
	s += "\nPress Enter to get install instructions."
	if bm.msg != "" {
		s += "\n\n" + bm.msg
	}
	s += "\n\nEsc · Back"
	return s
}

func styleStatus(status string, installed bool) string {
	if installed {
		return lipgloss.NewStyle().Foreground(green).Render(status)
	}
	return lipgloss.NewStyle().Foreground(accent).Render(status)
}
