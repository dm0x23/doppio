package sync

import (
	"fmt"
	"os"
	"strings"

	"github.com/dm0x23/doppio/internal/shell"
	"github.com/dm0x23/doppio/internal/storage"
)

func Run() error {
	shortcuts, err := storage.List()
	if err != nil {
		return fmt.Errorf("sync failed: %w", err)
	}

	adapters := []shell.Adapter{
		shell.BashAdapter{},
		shell.ZshAdapter{},
	}

	for _, adapter := range adapters {
		var block strings.Builder
		start, end := adapter.ManagedBlockMarkers()
		block.WriteString(start + "\n")
		for _, s := range shortcuts {
			block.WriteString(adapter.GenerateAlias(s.Name, s.Command) + "\n")
		}
		block.WriteString(end + "\n")

		rcPath := adapter.RCFilePath()
		content, err := os.ReadFile(rcPath)
		if err != nil {
			if os.IsNotExist(err) {
				os.WriteFile(rcPath, []byte(block.String()), 0o644)
				continue
			}
			return fmt.Errorf("sync failed: %w", err)
		}

		startIdx := strings.Index(string(content), start)
		endIdx := strings.Index(string(content), end)

		var newContent string
		if startIdx != -1 && endIdx != -1 {
			newContent = string(content[:startIdx]) + block.String() + string(content[endIdx+len(end)+1:])
		} else {
			newContent = string(content) + "\n" + block.String()
		}

		if err := os.WriteFile(rcPath, []byte(newContent), 0o644); err != nil {
			return fmt.Errorf("sync: failed to write %s: %w", rcPath, err)
		}
	}
	return nil
}
