package shell

import (
	"fmt"
	"os"
	"path/filepath"
)

type ZshAdapter struct{}

func (z ZshAdapter) GenerateAlias(name, command string) string {
	return fmt.Sprintf("alias %s='%s'", name, command)
}

func (z ZshAdapter) RCFilePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".zshrc")
}

func (z ZshAdapter) ManagedBlockMarkers() (string, string) {
	return "# >>> doppio managed >>>", "# <<< doppio managed <<<"
}
