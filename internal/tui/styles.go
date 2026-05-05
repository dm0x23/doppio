package tui

import "charm.land/lipgloss/v2"

var (
	// ANSI colour codes – adapt to terminal theme
	base      = lipgloss.Color("7")  // white (foreground)
	surface   = lipgloss.Color("0")  // black
	subtle    = lipgloss.Color("8")  // bright black (dark grey)
	highlight = lipgloss.Color("14") // bright cyan
	accent    = lipgloss.Color("9")  // bright red
	green     = lipgloss.Color("10") // bright green

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

	// Selection with a subtle dark overlay – not full opacity
	TableSelectedStyle = lipgloss.NewStyle().
				Background(subtle). // = ANSI 8 (dark grey)
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
