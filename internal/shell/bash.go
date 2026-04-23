package shell

import (
	"os"
	"path/filepath"
)

type BashAdapter struct{}

func (b BashAdapter) GenerateAlias(name, command string) string {
	return `alias ` + name + `='` + command + `'`
}

func (b BashAdapter) RCFilePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".bashrc")
}

func (b BashAdapter) ManagedBlockMarkers() (string, string) {
	return "# >>> doppio managed >>>", "# <<< doppio managed <<<"
}
