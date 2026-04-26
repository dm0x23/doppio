package watch

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/dm0x23/doppio/internal/storage"
	"github.com/dm0x23/doppio/internal/sync"
	"github.com/fsnotify/fsnotify"
)

func Start(dir string, shells []string, auto, includeExisting bool) error {
	if includeExisting {
		aliasExisting(dir, shells)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("failed to create watcher: %w", err)
	}
	defer watcher.Close()

	if err := watcher.Add(dir); err != nil {
		return fmt.Errorf("failed to watch directory: %w", err)
	}

	fmt.Printf("Watching %s for new folders... \n", dir)
	fmt.Println("Press CTRL-C to stop")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Create != 0 {
				handleNewItem(event.Name, shells, auto)
			}
			if event.Op&fsnotify.Remove != 0 {
				handleRemovedItem(event.Name)
			}

		case err := <-watcher.Errors:
			fmt.Printf("Watcher errors: %v\n", err)

		case <-signalChan:
			fmt.Println("watch stopped")
			return nil
		}
	}
}

func handleNewItem(path string, shells []string, auto bool) {
	info, err := os.Stat(path)
	if err != nil || !info.IsDir() {
		return
	}

	folderName := filepath.Base(path)
	suggestedName := SanitizeName(folderName)
	var aliasName string

	if auto {
		aliasName = suggestedName
	} else {
		name, skip := promptForAlias(folderName, suggestedName)
		if skip {
			return
		}
		aliasName = name
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("Failed to resolve path: %v\n", err)
		return
	}

	command := "cd " + absPath
	if err := storage.Add(aliasName, command, shells); err != nil {
		fmt.Printf("Failed to add alias: %v\n", err)
		return
	}

	if err := sync.Run(); err != nil {
		fmt.Printf("failed to sync: %v\n", err)
	}

	fmt.Printf("Alias '%s' -> %s added\n", aliasName, absPath)
}

func aliasExisting(dir string, shells []string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("could not read directory: %v\n", err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			name := SanitizeName(entry.Name())
			absPath, _ := filepath.Abs(entry.Name())
			command := "cd " + absPath

			if err := storage.Add(name, command, shells); err != nil {
				fmt.Printf("Skipping %s: %v\n", name, err)
				continue
			}
			fmt.Printf("Alias %s -> %s added\n", name, absPath)
		}
	}
	sync.Run()
}

func handleRemovedItem(path string) {
	folderName := filepath.Base(path)
	aliasName := SanitizeName(folderName)

	if err := storage.Remove(aliasName); err != nil {
		return
	}

	if err := sync.Run(); err != nil {
		fmt.Printf("Alias removed but sync failed: %v\n", err)
		return
	}

	fmt.Printf("Alias '%s' removed (folder deleted) \n", aliasName)
}
