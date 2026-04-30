package tui

import "charm.land/lipgloss/v2"

var (
	base      = lipgloss.Color("#f0f0f0")
	surface   = lipgloss.Color("#252525")
	subtle    = lipgloss.Color("#383838")
	highlight = lipgloss.Color("#7b61ff")
	accent    = lipgloss.Color("#ff5e5e")
	green     = lipgloss.Color("#00b894")

	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(highlight).
			PaddingLeft(1)

	TableHeaderStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#1a1a1a")).
				Background(highlight).
				Padding(0, 1)

	TableRowStyle = lipgloss.NewStyle().
			Padding(0, 1)

	TableSelectedStyle = lipgloss.NewStyle().
				Background(surface).
				Padding(0, 1).
				Foreground(highlight)

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
