package cmd

import (
	"fmt"

	"github.com/dm0x23/doppio/internal/sync"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Manually sync shortcuts to shell config files",
	RunE:  runSync,
}

func init() {
	rootCmd.AddCommand(syncCmd)
}

func runSync(cmd *cobra.Command, args []string) error {
	if err := sync.Run(); err != nil {
		return fmt.Errorf("Sync failed: %w", err)
	}
	fmt.Println("Synced to shell config ✓")
	return nil
}
