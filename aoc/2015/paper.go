package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput(file string) io.Reader {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	return f
}

func main() {

	total := 0
	ribbon := 0
	scanner := bufio.NewScanner(readInput("d2.input"))
	for scanner.Scan() {
		factors := strings.Split(scanner.Text(), "x")

		l, err := strconv.Atoi(factors[0])
		if err != nil {
			panic(err)
		}
		w, err := strconv.Atoi(factors[1])
		if err != nil {
			panic(err)
		}
		h, err := strconv.Atoi(factors[2])
		if err != nil {
			panic(err)
		}
		area1 := 2 * l * w
		area2 := 2 * l * h
		area3 := 2 * h * w

		minArea := area1
		if area2 < minArea {
			minArea = area2
		}
		if area3 < minArea {
			minArea = area3
		}
		total += area1 + area2 + area3 + minArea/2

		sides := []int{l, w, h}
		sort.Ints(sides)
		ribbon += l * w * h
		ribbon += sides[0]*2 + sides[1]*2

	}

	fmt.Println(total, ribbon)

}
