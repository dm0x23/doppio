package bootstrap

type Tool struct {
	Name        string
	Description string
	Brew        string // macOS
	Apt         string // Debian/Ubuntu
	Pacman      string // Arch
	Dnf         string // Fedora
	Zypper      string // OpenSUSE
	Apk         string // Alpine
}

var DefaultTools = []Tool{
	{
		Name:        "zoxide",
		Description: "Smarter cd command",
		Brew:        "zoxide",
		Apt:         "zoxide",
		Pacman:      "zoxide",
		Dnf:         "zoxide",
		Zypper:      "zoxide",
		Apk:         "zoxide",
	},
	{
		Name:        "fzf",
		Description: "Fuzzy finder",
		Brew:        "fzf",
		Apt:         "fzf",
		Pacman:      "fzf",
		Dnf:         "fzf",
		Zypper:      "fzf",
		Apk:         "fzf",
	},
	{
		Name:        "bat",
		Description: "Better cat with syntax highlighting",
		Brew:        "bat",
		Apt:         "bat",
		Pacman:      "bat",
		Dnf:         "bat",
		Zypper:      "bat",
		Apk:         "bat",
	},
	{
		Name:        "ripgrep",
		Description: "Faster grep",
		Brew:        "ripgrep",
		Apt:         "ripgrep",
		Pacman:      "ripgrep",
		Dnf:         "ripgrep",
		Zypper:      "ripgrep",
		Apk:         "ripgrep",
	},
	{
		Name:        "eza",
		Description: "Modern ls replacement",
		Brew:        "eza",
		Apt:         "eza",
		Pacman:      "eza",
		Dnf:         "eza",
		Zypper:      "eza",
		Apk:         "eza",
	},
}
