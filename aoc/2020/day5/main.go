package main

import (
	"fmt"
	"math"

	"github.com/Gnoale/adventofcode/puzzlein"
)

type seatCoor struct {
	row    int
	column int
	id     int
}

var bitTable = map[string]int{"F": 0, "B": 1, "R": 1, "L": 0}

func findMax(sc []seatCoor) int {
	m := 0
	for _, s := range sc {
		if s.id > m {
			m = s.id
		}
	}
	return m
}

func compSeat(seatm []string) []seatCoor {
	res := make([]seatCoor, 0)
	var r, c float64
	for _, s := range seatm {
		m := seatCoor{}
		r, c = 0, 0
		for i, b := range s {
			if bitTable[string(b)] == 1 {
				if i > 6 {
					// 2**2, 2**1, 2**0
					c += math.Pow(float64(2), float64(9-i))
				} else {
					// 2**6, 2**5 ... 2**0
					r += math.Pow(float64(2), float64(len(s)-4-i))
				}
			}
		}
		m.row = int(r)
		m.column = int(c)
		m.id = (m.row * 8) + m.column
		res = append(res, m)
	}
	return res
}

func findSeat(seats []seatCoor) {
	var p int
	for i, s := range seats {
		if i > 0 && (s.id-p > 1) {
			fmt.Println(s, seats[i-1])
		}
		p = s.id
	}

}

func main() {
	seatm, err := puzzlein.GetStr("./day5/input")
	if err != nil {
		panic(err)
	}
	// part1
	res := compSeat(seatm)
	max := findMax(res)
	fmt.Println(max)

	// part2
	seatid := func(s1, s2 *seatCoor) bool {
		return s1.id < s2.id
	}
	By(seatid).Sort(res)
	findSeat(res)

}
