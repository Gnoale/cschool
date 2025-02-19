package leetcode75

import (
	"strings"
)

func reverseWords(s string) string {
	splitted := strings.Split(strings.TrimSpace(s), " ")
	reversed := ""
	for i := len(splitted) - 1; i > 0; i-- {
		if splitted[i] != "" {
			reversed += splitted[i] + " "
		}
	}
	reversed += splitted[0]
	return reversed
}
