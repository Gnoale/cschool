package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func readInput(file string) string {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func main() {
	input := strings.Trim(readInput("d1.input"), "\n")

	steps := 0
	i := 1
	for len(input) > 0 {
		r, size := utf8.DecodeRuneInString(input)
		if r == '(' {
			steps += 1
		}
		if r == ')' {
			steps -= 1
		}
		input = input[size:]
		if steps == -1 {
			break
		}
		i += 1
	}

	fmt.Println(steps, i)

}
