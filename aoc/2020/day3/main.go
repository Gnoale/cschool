package main

import (
	"fmt"

	"github.com/Gnoale/adventofcode/puzzlein"
)

func countTree(m []string, x, y int) int {
	tc := 0
	j := x
	var s string
	for i := y; i < len(m); {
		L := len(m[i])
		for ; j < L && i < len(m); j += x {
			s = string(m[i][j])
			//		fmt.Println(s, i, j)
			if s == "#" {
				tc++
			}
			i += y
		}
		if j+y >= L {
			j -= L
		}
	}
	return tc

}

func main() {
	cart, err := puzzlein.GetStr("./day3/input")
	if err != nil {
		panic(err)
	}
	// part1
	tcb := countTree(cart, 3, 1)

	// part2
	tca := countTree(cart, 1, 1)
	tcc := countTree(cart, 5, 1)
	tcd := countTree(cart, 7, 1)
	tce := countTree(cart, 1, 2)
	fmt.Println(tca, tcb, tcc, tcd, tce)
	fmt.Println(tca * tcb * tcc * tcd * tce)
}
