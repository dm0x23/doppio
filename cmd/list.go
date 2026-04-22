package cmd

import (
	"fmt"

	"github.com/dm0x23/doppio/internal/storage"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all shortcuts",
	Long:  `Display all shortcuts currently managed by doppio`,
	RunE:  runList,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func runList(cmd *cobra.Command, args []string) error {
	shortcuts, err := storage.List()
	if err != nil {
		return fmt.Errorf("failed to load shortcuts: %w", err)
	}

	if len(shortcuts) == 0 {
		fmt.Println("No shortcuts added")
		fmt.Println("Add a shortcut using dop add <name> <command>")
	}

	fmt.Println("shortcuts")

	for _, s := range shortcuts {
		fmt.Printf("%s -> %s\n", s.Name, s.Command)
	}

	return nil
}
