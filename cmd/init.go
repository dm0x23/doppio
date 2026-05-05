package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Install Doppio's watch daemon as a user systemd service",
	Long:  `Creates and enables a systemd user service that runs 'dop watch' in the background.`,
	RunE:  runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInit(cmd *cobra.Command, args []string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("cannot find home directory: %w", err)
	}

	unitDir := filepath.Join(home, ".config", "systemd", "user")
	if err := os.MkdirAll(unitDir, 0o755); err != nil {
		return fmt.Errorf("failed to create systemd user dir: %w", err)
	}

	unitPath := filepath.Join(unitDir, "doppio-watch.service")
	serviceContent := fmt.Sprintf(`[Unit]
Description=Doppio Watch Daemon
After=network.target

[Service]
Type=simple
ExecStart=%s/go/bin/dop watch
Environment=DBUS_SESSION_BUS_ADDRESS=unix:path=/run/user/%U/bus
Restart=on-failure
RestartSec=10

[Install]
WantedBy=default.target
`, home)

	if err := os.WriteFile(unitPath, []byte(serviceContent), 0o644); err != nil {
		return fmt.Errorf("failed to write service file: %w", err)
	}
	fmt.Println("Created systemd service:", unitPath)

	if err := exec.Command("systemctl", "--user", "daemon-reload").Run(); err != nil {
		return fmt.Errorf("failed to reload systemd: %w", err)
	}
	fmt.Println("Reloaded systemd user daemon")

	if err := exec.Command("systemctl", "--user", "enable", "--now", "doppio-watch.service").Run(); err != nil {
		return fmt.Errorf("failed to enable service: %w", err)
	}
	fmt.Println("Service enabled and started!")

	fmt.Println("\nYou can manage the daemon with:")
	fmt.Println("  systemctl --user status doppio-watch")
	fmt.Println("  systemctl --user stop doppio-watch")
	fmt.Println("  systemctl --user disable doppio-watch")
	return nil
}
