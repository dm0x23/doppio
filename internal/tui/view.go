package tui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m Model) View() tea.View {
	if m.width == 0 {
		return tea.NewView("Loading...")
	}

	header := HeaderStyle.Render("☕ Doppio")

	var body string
	switch m.currentScreen() {
	case "list":
		body = m.renderTable()
	case "add":
		body = m.addForm.View()
	case "delete-confirm":
		body = m.deleteConfirm.View()
	case "watch":
		body = m.watchModel.View()
	case "bootstrap":
		body = m.bootstrapModel.View()
	case "bulk-delete-confirm":
		body = m.bulkDeleteConfirm.View()
	default:
		body = "Unknown screen"
	}

	if m.currentScreen() != "list" {
		body = lipgloss.Place(m.width, m.height-2,
			lipgloss.Center, lipgloss.Center,
			ModalStyle.Render(body),
		)
	}

	help := "(a)dd (d)elete (w)atch (b)ootstrap (s)ync"
	status := ""
	if m.statusMsg != "" {
		status = lipgloss.NewStyle().Foreground(green).Render(m.statusMsg)
	}
	footer := lipgloss.JoinHorizontal(lipgloss.Top,
		HelpStyle.Render(help),
		lipgloss.NewStyle().Width(m.width-lipgloss.Width(help)-lipgloss.Width(status)).Render(""),
		status,
	)
	footer = StatusBarStyle.Width(m.width).Render(footer)

	content := lipgloss.JoinVertical(lipgloss.Left,
		header,
		body,
		footer,
	)
	return tea.NewView(content)
}

func (m Model) renderTable() string {
	header := TableHeaderStyle.Render("  Shortcut") + "  " + TableHeaderStyle.Render("Command")
	rows := []string{header}
	for i, sc := range m.shortcuts {
		marker := " "
		if sc.Selected {
			marker = "✓"
		}
		cursor := marker + " "
		if m.cursor == i {
			cursor = CursorMarker.Render(">") + marker + " "
		}

		row := lipgloss.JoinHorizontal(lipgloss.Top,
			TableRowStyle.Render(fmt.Sprintf("%s %s", cursor, sc.Shortcut.Name)),
			TableRowStyle.Render(sc.Shortcut.Command),
		)
		if m.cursor == i {
			row = TableSelectedStyle.Render(row)
		}
		rows = append(rows, row)
	}
	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
