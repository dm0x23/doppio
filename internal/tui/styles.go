package tui

import "charm.land/lipgloss/v2"

var (
	base      = lipgloss.Color("7")
	surface   = lipgloss.Color("0") // black
	subtle    = lipgloss.Color("8") // bright black (dark grey)
	highlight = lipgloss.Color("12")
	accent    = lipgloss.Color("1")
	green     = lipgloss.Color("2")

	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(highlight).
			PaddingLeft(1)

	TableHeaderStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(base).
				Background(highlight).
				Padding(0, 1)

	TableRowStyle = lipgloss.NewStyle().
			Padding(0, 1)

	TableSelectedStyle = lipgloss.NewStyle().
				Background(subtle).
				Padding(0, 1).
				Foreground(base)

	CursorMarker = lipgloss.NewStyle().
			Foreground(highlight).
			Bold(true)

	StatusBarStyle = lipgloss.NewStyle().
			Background(surface).
			Foreground(subtle).
			Padding(0, 1)

	HelpStyle = lipgloss.NewStyle().
			Foreground(subtle)

	ModalStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(highlight).
			Padding(2)

	InputStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			Padding(0, 1)
)
