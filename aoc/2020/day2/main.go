package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/Gnoale/adventofcode/puzzlein"
)

func findPart1Comply(plist []string) int {
	var count int
	for _, line := range plist {
		l := strings.Split(line, " ")
		min, _ := strconv.Atoi(strings.Split(l[0], "-")[0])
		max, _ := strconv.Atoi(strings.Split(l[0], "-")[1])
		letter, _ := utf8.DecodeRuneInString(strings.Trim(l[1], ":"))
		for _, r := range l[2] {
			if letter == r {
				min -= 1
				max -= 1
			}
		}
		if min <= 0 && max >= 0 {
			count += 1
		}
	}
	return count
}

func findPart2Comply(plist []string) int {
	var count int
	for _, line := range plist {
		l := strings.Split(line, " ")
		fst, _ := strconv.Atoi(strings.Split(l[0], "-")[0])
		sec, _ := strconv.Atoi(strings.Split(l[0], "-")[1])
		letter, _ := utf8.DecodeRuneInString(strings.Trim(l[1], ":"))
		m := 0
		for i, r := range l[2] {
			if letter == r {
				if i+1 == fst || i+1 == sec {
					m += 1
				}
			}
		}
		if m == 1 {
			count += 1
		}
	}
	return count
}

func main() {
	pass, err := puzzlein.GetStr("./day2/input")
	if err != nil {
		panic(err)
	}
	c := findPart1Comply(pass)
	fmt.Println(c)

	c = findPart2Comply(pass)
	fmt.Println(c)

}
