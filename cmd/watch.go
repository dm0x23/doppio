package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dm0x23/doppio/internal/watch"
	"github.com/fsnotify/fsnotify"
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
	var dirsToWatch []string

	if len(args) > 0 {
		newDir := args[0]
		info, err := os.Stat(newDir)
		if err != nil {
			return fmt.Errorf("cannot access directory: %w", err)
		}
		if !info.IsDir() {
			return fmt.Errorf("not a directory: %s", newDir)
		}

		existingDirs, err := watch.LoadDirs()
		if err != nil {
			return fmt.Errorf("failed to read watch config: %w", err)
		}

		alreadyWatched := false
		for _, d := range existingDirs {
			if d == newDir {
				alreadyWatched = true
				break
			}
		}
		if !alreadyWatched {
			existingDirs = append(existingDirs, newDir)
			if err := watch.SaveDirs(existingDirs); err != nil {
				return fmt.Errorf("failed to save watch config: %w", err)
			}
			fmt.Printf("Added %s to watched directories\n", newDir)
		}

		dirsToWatch = existingDirs
	} else {
		var err error
		dirsToWatch, err = watch.LoadDirs()
		if err != nil {
			return fmt.Errorf("failed to load watch config: %w", err)
		}
		if len(dirsToWatch) == 0 {
			fmt.Println("No watched directories. Add one with: dop watch <path>")
			return nil
		}
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("failed to create watcher: %w", err)
	}
	defer watcher.Close()

	for _, dir := range dirsToWatch {
		if err := watcher.Add(dir); err != nil {
			fmt.Printf("Could not watch %s: %v\n", dir, err)
		} else {
			fmt.Printf("Watching %s\n", dir)
		}
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Press Ctrl+C to stop all watchers.")

	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Create != 0 {
				watch.HandleNewItem(event.Name, shells, auto)
			}

		case err := <-watcher.Errors:
			fmt.Printf("Watch error: %v\n", err)

		case <-signalChan:
			fmt.Println("\nWatch stopped.")
			return nil
		}
	}
}
