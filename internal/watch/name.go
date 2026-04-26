package watch

import (
	"strings"
	"unicode"
)

func SanitizeName(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "-")

	var result strings.Builder
	for _, r := range name {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '_' {
			result.WriteRune(r)
		}
	}

	sanitized := result.String()
	if sanitized == "" {
		return name
	}
	return sanitized
}
