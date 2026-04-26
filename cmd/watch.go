package cmd

import (
	"fmt"
	"os"

	"github.com/dm0x23/doppio/internal/watch"
	"github.com/spf13/cobra"
)

var (
	auto            bool
	includeExisting bool
)

var watchCmd = &cobra.Command{
	Use:   "watch <directory>",
	Short: "Watch a directory and autoalias new folders",
	Long: `Watch monitors a directory for new subfolders and 
prompts you to create aliases for them automatically

Use --auto to skip prompts and create aliases for them automatically,
Use --include-existing to alias folders already in directory
	`,
	Args: cobra.ExactArgs(1),
	RunE: runWatch,
}

func init() {
	watchCmd.Flags().BoolVar(&auto, "auto", false, "Skip confirmation and auto-create aliases")
	watchCmd.Flags().BoolVar(&includeExisting, "include-existing", false, "Also create aliases for existing folders")

	watchCmd.Flags().StringSliceVar(&shells, "shells", []string{"zsh", "bash"}, "Target shells")
	rootCmd.AddCommand(watchCmd)
}

func runWatch(cmd *cobra.Command, args []string) error {
	dir := args[0]

	info, err := os.Stat(dir)
	if err != nil {
		return fmt.Errorf("cannot access directory %w", err)
	}

	if !info.IsDir() {
		return fmt.Errorf("not a directory %w", err)
	}

	return watch.Start(dir, shells, auto, includeExisting)
}
