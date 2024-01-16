package utils

import (
	"io"
	"os"
)

func ReadInput(file string) io.Reader {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	return f
}

func ReadInputString(file string) string {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return string(b)
}
