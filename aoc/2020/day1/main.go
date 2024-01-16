package main

import (
	"fmt"

	"github.com/Gnoale/adventofcode/puzzlein"
)

func findProductN(s []int) int {
	for i, t := range s {
		for j := i; s[j] < (2020-t) && j < len(s)-1; j++ {
			u := s[j]
			for k := j; s[k] <= (2020-t-u) && k < len(s)-1; k++ {
				if (t + u + s[k]) == 2020 {
					return (t * u * s[k])
				}
			}
		}
	}
	return 0
}

func findProductNr(s []int) int {
	l := len(s)
	if l == 1 {
		return 0
	}
	L := s[:l/2]
	R := s[l/2:]
	if r := findProductN(L); r != 0 {
		return r
	}
	if r := findProductN(R); r != 0 {
		return r
	}
	findProductNr(L)
	findProductNr(R)
	return 0
}

func findProductNn(s []int) int {
	for i, t := range s {
		for j := i; j < len(s)-1; j++ {
			u := s[j]
			for k := j; k < len(s)-1; k++ {
				if (t + u + s[k]) == 2020 {
					return (t * u * s[k])
				}
			}
		}
	}
	return 0
}

func main() {

	codes, err := puzzlein.GetInt("./day1/input")
	if err != nil {
		panic(err)
	}

	//r := findProductN(codes)
	//fmt.Println(r)

	r := findProductNr(codes)
	fmt.Println(r)

}
