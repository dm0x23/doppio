# ☕ Doppio

> A double shot of speed for your shell.

`dop` is a shortcut manager for your terminal. Create aliases, watch directories for new projects, and bootstrap your development environment — all from the command line.

--- 
## Demo

<img width="1200" height="600" alt="doppio-demo" src="https://github.com/user-attachments/assets/5731c7de-0073-41a8-a5a3-d3628eeb0539" />
> Demo for the CLI version of doppio

https://github.com/user-attachments/assets/d27b75e7-d6b0-4651-a0b0-989bb4c0084c
> Demo for the TUI version of doppio

---

## Features

- **Add shortcuts:** `dop add proj "cd ~/Desktop/projects"`
- **List them:** `dop list`
- **Remove them:** `dop remove proj docs gs`
- **Auto-sync:** Shell config updated automatically after every change
- **Watch mode:** Point at a directory and auto-alias new projects as they're created
- **Bootstrap:** One command to install favorite CLI tools and configure sensible defaults
- **Managed blocks:** Never touches your other aliases, everything lives between markers
- **TUI:** Interactive terminal interface for managing shortcuts
- **Background daemon:** dop watch runs as a systemd service, surviving reboots
- **Interactive naming:** Desktop notification + terminal pop‑up to name new folders
- **TUI with multi‑select:** Space to select, bulk delete, Watch & Bootstrap screens
---

## How it works
<img width="2084" height="1922" alt="doppio" src="https://github.com/user-attachments/assets/c0b0ad8d-58ff-43c2-a3e9-1cd849fa5889" />

1. You type `dop add <name> <command>`
2. Doppio stores it in ~/.config/doppio/shortcuts.json
3. The sync engine writes the alias to a managed block in your shell config
4. You type the shortcut name and it just works
5. Or speed it up further by setting up a watched directory with `dop watch <path>` and every new folder there is made alias automatically

--- 
## Installation

### Via Go Install

```bash
go install github.com/dm0x23/doppio/cmd/dop@latest
```

Make sure `$HOME/go/bin` is in your PATH:
```bash
export PATH="$HOME/go/bin:$PATH"
```

### From Source
```bash
git clone https://github.com/dm0x23/doppio.git
cd doppio
make build
make install
```

> run `dop init` if you want to add it as a background-service for dop watch. Feel free to configure the path if you have moved the binary elsewhere: `systemctl --user edit --full doppio-watch`
---

## Quickstart
```bash
# Add a shortcut
dop add proj "cd ~/Desktop/projects"

# Source your shell
source ~/.zshrc

# Use it
proj     # → jumps to ~/Desktop/projects

# List all shortcuts
dop list

# Remove one or more
dop remove proj
```
---

## File locations
| What | Location |
|------|----------|
| Shortcut data | `~/.config/doppio/shortcuts.json` |
| Watch config | `~/.config/doppio/watch.json` |
| Shell managed block | `~/.zshrc`, `~/.bashrc` (between `# >>> doppio managed >>>` markers) |
| Installed binary | `~/go/bin/dop` |
---

## Commands

| Command | Description |
|---------|-------------|
| `dop add <name> <command>` | Add a new shortcut |
| `dop list` | List all shortcuts |
| `dop remove <name>` | Remove a shortcut |
| `dop sync` | Manually sync shortcuts to shell configs |
| `dop watch <directory>` | Watch a directory and auto-alias new folders |
| `dop bootstrap` | Install recommended CLI tools + configure aliases |
| `dop completion <shell>` | Generate shell autocompletion script |
| `dop --version` | Print version information |
| `dop remove <name> [<name> ...]` | Remove one or more shortcuts (bulk) |
| `dop tui`                         | Launch the interactive terminal UI |
| `dop init`                        | Install the watch daemon as a user systemd service |
---

## Roadmap
- [x] Manual shortcut management
- [x] Watch mode
- [x] Bootstrap: The one command development setup
- [x] Make watch into something that runs in the background
- [x] TUI with BubbleTea
- [ ] Support for fish, powershell etc.

## Built with
- Go
- Cobra - CLI Framework
- BubbleTea - Framework for TUI
- LipGloss - A way to style the BubbleTea TUI
- Bubbles - Components for BubbleTea
- fsnotify - File system watcher

## Author
_Denisha|denisha.co.in_
