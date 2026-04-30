package cmd

import (
	tea "charm.land/bubbletea/v2"
	"github.com/dm0x23/doppio/internal/tui"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Launch the Doppio TUI",
	RunE: func(cmd *cobra.Command, args []string) error {
		program := tea.NewProgram(tui.NewModel())
		_, err := program.Run()
		return err
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
