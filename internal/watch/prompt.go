package watch

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func promptForAlias(folderName, suggestedName string) (string, bool) {
	fmt.Printf("\nNew folder detected: %s\n", folderName)
	fmt.Printf("Alias name [%s]: ", suggestedName)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "cancel" || input == "no" || input == "n" {
		fmt.Println("Skipped")
		return "", true
	}

	if input == "" {
		return suggestedName, false
	}

	return SanitizeName(input), false
}
