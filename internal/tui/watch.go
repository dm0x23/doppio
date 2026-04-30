package tui

import tea "charm.land/bubbletea/v2"

type WatchModel struct{}

func NewWatchModel() WatchModel                   { return WatchModel{} }
func (wm *WatchModel) Update(msg tea.Msg) tea.Cmd { return nil }
func (wm WatchModel) View() string {
	return "👀 Watch Mode\n\nWatch directories for new projects and auto‑alias them.\n\nPress Esc to return."
}
