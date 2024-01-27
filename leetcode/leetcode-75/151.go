package main

import (
	"fmt"
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

func main() {
	fmt.Println("blue is sky the" == reverseWords("the sky is blue"))
	fmt.Println("world hello" == reverseWords("  hello world  "))
	fmt.Println("example good a" == reverseWords("a   good example"))
}
