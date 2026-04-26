package bootstrap

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type PackageManager struct {
	Name     string
	Install  string
	Sudo     bool
	CheckCmd string
}

func DetectPackageManager() (*PackageManager, error) {
	managers := []PackageManager{
		{Name: "brew", Install: "brew install", Sudo: false, CheckCmd: "brew"},
		{Name: "apt", Install: "apt install -y", Sudo: true, CheckCmd: "apt"},
		{Name: "pacman", Install: "pacman -S --needed --noconfirm", Sudo: true, CheckCmd: "pacman"},
		{Name: "dnf", Install: "dnf install -y", Sudo: true, CheckCmd: "dnf"},
		{Name: "zypper", Install: "zypper install -y", Sudo: true, CheckCmd: "zypper"},
		{Name: "apk", Install: "apk add", Sudo: true, CheckCmd: "apk"},
	}

	for _, pm := range managers {
		if _, err := exec.LookPath(pm.CheckCmd); err == nil {
			return &pm, nil
		}
	}

	return nil, fmt.Errorf("no supported package managers found")
}

func InstallTools(pm *PackageManager, tools []Tool) error {
	for _, tool := range tools {
		pkgName := getPackageName(tool, pm.Name)
		if pkgName == "" {
			fmt.Printf("Skipping %s: not available for %s\n", tool.Name, pm.Name)
			continue
		}
		installArgs := strings.Fields(pm.Install)

		var cmd *exec.Cmd
		if pm.Sudo {
			allArgs := append(installArgs, pkgName)
			cmd = exec.Command("sudo", allArgs...)
		} else {
			allArgs := append(installArgs[1:], pkgName)
			cmd = exec.Command(installArgs[0], allArgs...)
		}

		fmt.Printf("Installing %s...\n", tool.Name)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Printf("Failed to install %s: %v\n", tool.Name, err)
		} else {
			fmt.Printf("%s installed\n", tool.Name)
		}
	}
	return nil
}

func getPackageName(tool Tool, pmName string) string {
	switch pmName {
	case "brew":
		return tool.Brew
	case "apt":
		return tool.Apt
	case "pacman":
		return tool.Pacman
	case "dnf":
		return tool.Dnf
	case "zypper":
		return tool.Zypper
	case "apk":
		return tool.Apk
	default:
		return ""
	}
}
