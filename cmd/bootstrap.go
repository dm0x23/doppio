package cmd

import (
	"fmt"

	"github.com/dm0x23/doppio/internal/bootstrap"
	"github.com/dm0x23/doppio/internal/storage"
	"github.com/dm0x23/doppio/internal/sync"
	"github.com/spf13/cobra"
)

var dryRun bool

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstrap your development environment",
	Long:  `Installs recommended CLI tools and configures sensible defaults.`,
	RunE:  runBootstrap,
}

func init() {
	bootstrapCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Show what all would be installed without actually installing it")
	rootCmd.AddCommand(bootstrapCmd)
}

func runBootstrap(cmd *cobra.Command, args []string) error {
	pm, err := bootstrap.DetectPackageManager()
	if err != nil {
		return err
	}

	fmt.Printf("Detected package manager %s\n", pm.Name)

	if dryRun {
		fmt.Println("\nWould Install:")
		for _, tool := range bootstrap.DefaultTools {
			fmt.Printf(" - %s (%s)\n", tool.Name, tool.Description)
		}
		fmt.Println("\nWould configure aliases:")
		for _, alias := range bootstrap.RecommendedAliases {
			fmt.Printf(" - alias %s = '%s'\n", alias.Name, alias.Command)
		}
		return nil
	}

	if err := bootstrap.InstallTools(pm, bootstrap.DefaultTools); err != nil {
		return err
	}

	fmt.Println("\nConfiguring aliases...")
	for _, alias := range bootstrap.RecommendedAliases {
		if err := storage.Add(alias.Name, alias.Command, shells); err != nil {
			fmt.Printf("Skipping alias '%s': %v\n", alias.Name, err)
		} else {
			fmt.Printf("alias %s='%s' added\n", alias.Name, alias.Command)
		}
	}

	if err := sync.Run(); err != nil {
		return fmt.Errorf("bootstrap completed but sync failed: %w", err)
	}

	fmt.Println("\n Bootstrap complete! Run 'source ~/.zshrc' to apply changes.")
	return nil
}
