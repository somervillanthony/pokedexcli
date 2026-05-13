package repl

import (
	"strings"
)

func cleanInput(text string) []string {
	lwrText := strings.ToLower(text)
	return strings.Fields(lwrText)
}
