package aoc

import (
	"2023/pkg/utils"
	"bufio"
	"fmt"
	"strconv"
	"unicode/utf8"
)

func Snow() {

	total := 0
	scanner := bufio.NewScanner(utils.ReadInput("d1.input"))
	for scanner.Scan() {

		line := scanner.Text()

		first, last := 0, 0

		var err error
		var num int

		fmt.Println(line)

		for len(line) > 0 {
			// runtime.Breakpoint()

			r, n := utf8.DecodeRuneInString(line)

			// try to convert to int
			num, err = strconv.Atoi(string(r))
			if err == nil {
				if first == 0 {
					first = num
				}
				last = num
			}

			// part 2 we need to look forward if the number is written as as string
			if len(line) >= 3 {
				switch {
				case line[:3] == "one":
					if first == 0 {
						first = 1
					}
					last = 1
					line = line[2:]
					continue
				case line[:3] == "two":
					if first == 0 {
						first = 2
					}
					last = 2
					line = line[2:]
					continue
				case line[:3] == "six":
					if first == 0 {
						first = 6
					}
					last = 6
					line = line[2:]
					continue
				}
			}
			if len(line) >= 4 {
				switch {
				case line[:4] == "nine":
					if first == 0 {
						first = 9
					}
					last = 9
					line = line[3:]
					continue
				case line[:4] == "four":
					if first == 0 {
						first = 4
					}
					last = 4
					line = line[3:]
					continue
				case line[:4] == "five":
					if first == 0 {
						first = 5
					}
					last = 5
					line = line[3:]
					continue
				}
			}
			if len(line) >= 5 {
				switch {
				case line[:5] == "three":
					if first == 0 {
						first = 3
					}
					last = 3
					line = line[4:]
					continue
				case line[:5] == "seven":
					if first == 0 {
						first = 7
					}
					last = 7
					line = line[4:]
					continue
				case line[:5] == "eight":
					if first == 0 {
						first = 8
					}
					last = 8
					line = line[4:]
					continue
				}
			}
			line = line[n:]
		}

		c := first*10 + last
		fmt.Println(first, last, c)
		total += c
	}

	fmt.Println(total)

}
