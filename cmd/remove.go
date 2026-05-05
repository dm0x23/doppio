package cmd

import (
	"fmt"
	"strings"

	"github.com/dm0x23/doppio/internal/storage"
	"github.com/dm0x23/doppio/internal/sync"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [name]",
	Short: "Remove a shortcut",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	RunE:  runRemove,
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

func runRemove(cmd *cobra.Command, args []string) error {
	var failed []string
	for _, name := range args {
		if err := storage.Remove(name); err != nil {
			failed = append(failed, name)
		}
	}

	if err := sync.Run(); err != nil {
		return fmt.Errorf("sync failed: %w", err)
	}

	removed := len(args) - len(failed)
	if removed > 0 {
		fmt.Printf("Removed %d shortcut(s)\n", removed)
	}
	if len(failed) > 0 {
		fmt.Printf("Not found: %s\n", strings.Join(failed, ", "))
	}
	return nil
}
