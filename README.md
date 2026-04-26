# ☕ Doppio

> A double shot of speed for your shell.

`dop` is a shortcut manager for your terminal. Create aliases, watch directories for new projects, and bootstrap your development environment — all from the command line.

--- 
## Demo

https://github.com/user-attachments/assets/1d0a1159-3a1b-47f7-b119-76121963d7de

---

## Features

- **Add shortcuts:** `dop add proj "cd ~/Desktop/projects"`
- **List them:** `dop list`
- **Remove them:** `dop remove proj`
- **Auto-sync:** Shell config updated automatically after every change
- **Watch mode:** Point at a directory and auto-alias new projects as they're created
- **Bootstrap:** One command to install favorite CLI tools and configure sensible defaults
- **Managed blocks:** Never touches your other aliases, everything lives between markers
- **TUI:** Interactive terminal interface for managing shortcuts *(coming soon)*

---

## How it works
<img width="2337" height="748" alt="dop v0 pic" src="https://github.com/user-attachments/assets/67baabde-1428-47ee-a861-342de6dfeb4b" />

> Updated architecture coming soon

1. You type `dop add <name> <command>`
2. Doppio stores it in ~/.config/doppio/shortcuts.json
3. The sync engine writes the alias to a managed block in your shell config
4. You type the shortcut name and it just works

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
make install
```
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

# Remove one
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
---

## Roadmap
- [x] Manual shortcut management
- [x] Watch mode
- [x] Bootstrap: The one command development setup
- [ ] Make watch into something that runs in the background
- [ ] TUI with BubbleTea
- [ ] Support for fish, powershell etc.

## Built with
- Go
- Cobra - CLI Framework
- fsnotify - File system watcher

## Author
_Denisha|denisha.co.in_
