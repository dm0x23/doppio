package tui

import tea "charm.land/bubbletea/v2"

type BootstrapModel struct{}

func NewBootstrapModel() BootstrapModel               { return BootstrapModel{} }
func (bm *BootstrapModel) Update(msg tea.Msg) tea.Cmd { return nil }
func (bm BootstrapModel) View() string {
	return "🛠️ Bootstrap\n\nInstall recommended CLI tools and set up aliases.\n\n(Coming soon — use the CLI for now)\n\nPress Esc to return."
}
