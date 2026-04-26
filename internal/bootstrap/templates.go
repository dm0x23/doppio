package bootstrap

import "os/exec"

var RecommendedAliases = []struct {
	Name    string
	Command string
}{
	{Name: "ls", Command: "eza -a --icons"},
	{Name: "ll", Command: "eza -la --icons"},
	{Name: "tree", Command: "eza --icons --tree"},
	{Name: "bat", Command: getBatCommand()},
	{Name: "cd", Command: "z"},
	{Name: "grep", Command: "rg"},
}

func getBatCommand() string {
	if _, err := exec.LookPath("batcat"); err == nil {
		return "batcat"
	}
	return "bat"
}
