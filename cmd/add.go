package cmd

import (
	"fmt"

	"github.com/dm0x23/doppio/internal/storage"
	"github.com/dm0x23/doppio/internal/sync"
	"github.com/spf13/cobra"
)

var shells []string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   `add [name] [command]`,
	Short: `add a new shortcut`,
	Long:  ``,
	Args:  cobra.ExactArgs(2),
	RunE:  runAdd,
}

func init() {
	addCmd.Flags().StringSliceVarP(&shells, "shells", "s", []string{"zsh", "bash"}, "Target shells (comma-separated)")
	rootCmd.AddCommand(addCmd)
}

func runAdd(cmd *cobra.Command, args []string) error {
	name := args[0]
	command := args[1]

	if err := storage.Add(name, command, shells); err != nil {
		return fmt.Errorf("failed to add shortcut %w", err)
	}

	fmt.Printf("Shortcut %s add\n", name)
	if err := sync.Run(); err != nil {
		return fmt.Errorf("shortcut added but sync failed: %w", err)
	}

	fmt.Println("Synced to shell config ✓")
	return nil
}
