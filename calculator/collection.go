package calculator

import (
	// "fmt"
	"strings"
)

func TextContainsMaps(str string) map[string]string {
	words := strings.Split(str, "")
	mk := make(map[string]string)
	for _, v := range words {
		mk[v] = v
	}
	return mk
}

func TextContains(str string) []string {
	words := strings.Split(str, "")
	return words
}