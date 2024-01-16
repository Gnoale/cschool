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
	input := strings.Trim(readInput("d3.input"), "\n")

	n, s, e, w := 0, 0, 0, 0

	//grid := make([][]bool, 100)
	/*
		[t,t,
		[
		[
		..
	*/
	rows, columns := 0, 0
	for len(input) > 0 {
		r, size := utf8.DecodeRuneInString(input)
		if r == '^' {
			n += 1
			rows += 1
		}
		if r == '>' {
			e += 1
			columns += 1
		}
		if r == 'v' {
			s += 1
			rows -= 1
		}
		if r == '<' {
			w += 1
			columns -= 1
		}
		input = input[size:]
	}
	fmt.Println("rows =", rows, "columns =", columns)
	// rows = 67 columns = -65
	grid := make([][]bool, 200)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]bool, 200)
	}

	i := len(grid) / 2
	j := i

	input = strings.Trim(readInput("d3.input"), "\n")

	for len(input) > 0 {
		r, size := utf8.DecodeRuneInString(input)
		if r == '^' {
			i -= 1
			grid[i][j] = true
		}
		if r == '>' {
			j += 1
			grid[i][j] = true
		}
		if r == 'v' {
			i += 1
			grid[i][j] = true
		}
		if r == '<' {
			j -= 1
			grid[i][j] = true
		}
		input = input[size:]
	}

	houses := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] {
				houses += 1
			}
		}
	}
	fmt.Println(houses)

	// part2
	// Santa and robot santa
	input = strings.Trim(readInput("d3.input"), "\n")
	grids := make([][]bool, 200)
	for i := 0; i < len(grids); i++ {
		grids[i] = make([]bool, 200)
	}
	gridr := make([][]bool, 200)
	for i := 0; i < len(gridr); i++ {
		gridr[i] = make([]bool, 200)
	}

	i = len(grid) / 2
	j = i
	k, l := i, j
	grids[i][j] = true // santa only deliver a present at start

	x := 0
	for len(input) > 0 {
		r, size := utf8.DecodeRuneInString(input)
		if r == '^' {
			if x%2 == 0 {
				i -= 1
				grids[i][j] = true
			} else {
				k -= 1
				gridr[k][l] = true
			}
		}
		if r == '>' {
			if x%2 == 0 {
				j += 1
				grids[i][j] = true
			} else {
				l += 1
				gridr[k][l] = true
			}
		}
		if r == 'v' {
			if x%2 == 0 {
				i += 1
				grids[i][j] = true
			} else {
				k += 1
				gridr[k][l] = true
			}
		}
		if r == '<' {
			if x%2 == 0 {
				j -= 1
				grids[i][j] = true
			} else {
				l -= 1
				gridr[k][l] = true
			}
		}
		input = input[size:]
		x += 1
	}

	houses = 0
	for i := 0; i < len(grids); i++ {
		for j := 0; j < len(grids[i]); j++ {
			if grids[i][j] || gridr[i][j] {
				houses += 1
			}
		}
	}
	fmt.Println(houses)

}
