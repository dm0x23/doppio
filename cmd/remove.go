package cmd

import (
	"fmt"

	"github.com/dm0x23/doppio/internal/storage"
	"github.com/dm0x23/doppio/internal/sync"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove [name]",
	Short: "Remove a shortcut",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	RunE:  runRemove,
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

func runRemove(cmd *cobra.Command, args []string) error {
	name := args[0]
	if err := storage.Remove(name); err != nil {
		return fmt.Errorf("failed to remove shortcut %w", err)
	}

	fmt.Printf("Shortcut %s removed\n", name)
	if err := sync.Run(); err != nil {
		return fmt.Errorf("shortcut removed but sync failed: %w", err)
	}

	fmt.Println("Synced to shell config ✓")
	return nil
}
