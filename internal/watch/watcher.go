package watch

import (
	"fmt"
	"os"
	"os/exec"
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
				HandleNewItem(event.Name, shells, auto)
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

func HandleNewItem(path string, shells []string, auto bool) {
	info, err := os.Stat(path)
	if err != nil || !info.IsDir() {
		return
	}

	folderName := filepath.Base(path)
	suggestedName := SanitizeName(folderName)

	if auto {
		absPath, _ := filepath.Abs(path)
		command := "cd " + absPath
		if err := storage.Add(suggestedName, command, shells); err != nil {
			fmt.Printf("Failed to add alias: %v\n", err)
			return
		}
		if err := sync.Run(); err != nil {
			fmt.Printf("Alias added but sync failed: %v\n", err)
			return
		}
		fmt.Printf("Alias '%s' → %s\n", suggestedName, absPath)
		return
	}

	notifySend(
		"Doppio Watch",
		fmt.Sprintf("New folder: %s\nSuggested alias: %s\nRun in terminal: dop add %s \"cd %s\"",
			folderName, suggestedName, suggestedName, path),
	)

	exePath, err := os.Executable()
	if err != nil {
		exePath = "dop"
	}

	promptAlias(folderName, suggestedName, path, exePath)
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

func notifySend(title, message string) {
	cmd := exec.Command("notify-send", title, message, "--icon=terminal")
	cmd.Run()
}

func promptAlias(folderName, suggestedName, fullPath string, dopPath string) {
	script := fmt.Sprintf(
		`echo "New folder: %s"
echo "Suggested alias: %s"
read -p "Enter alias name (or Enter for '%s', or 'cancel' to skip): " alias
if [ -z "$alias" ]; then
    alias="%s"
fi
if [ "$alias" != "cancel" ]; then
   	%s add "$alias" "cd %s"
    echo "Alias added. Press Enter to close."
    read
fi`,
		folderName,
		suggestedName,
		suggestedName,
		suggestedName,
		dopPath,
		fullPath,
	)

	if _, err := exec.LookPath("x-terminal-emulator"); err == nil {
		cmd := exec.Command("x-terminal-emulator", "-e", "bash", "-c", script)
		cmd.Start()
		return
	}

	terminals := []struct {
		name string
		args []string
	}{
		{"kitty", []string{"-e", "bash", "-c", script}},
		{"alacritty", []string{"-e", "bash", "-c", script}},
		{"gnome-terminal", []string{"--", "bash", "-c", script}},
		{"konsole", []string{"-e", "bash", "-c", script}},
		{"xterm", []string{"-e", "bash", "-c", script}},
	}

	for _, term := range terminals {
		if _, err := exec.LookPath(term.name); err == nil {
			cmd := exec.Command(term.name, term.args...)
			cmd.Start()
			return
		}
	}
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
