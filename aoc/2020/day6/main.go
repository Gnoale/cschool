package main

import (
	"fmt"

	"github.com/Gnoale/adventofcode/puzzlein"
)

func findMatchAny(groups [][]string) {
	t := 0
	for _, group := range groups {
		m := make(map[string]int)
		for _, answers := range group {
			for _, a := range answers {
				if _, found := m[string(a)]; !found {
					m[string(a)] = 1
					t++
				}
			}
		}
		fmt.Println(m)
	}
	fmt.Println(t)
}

// part2
func findMatchEvery(groups [][]string) {
	t := 0
	for _, group := range groups {
		m := make(map[string]int)
		for _, answers := range group {
			for _, a := range answers {
				m[string(a)]++
			}
		}
		for _, v := range m {
			if v == len(group) {
				t++
			}
		}
		fmt.Println(m)
	}
	fmt.Println(t)
}

func main() {
	groups, err := puzzlein.GetStrEl("./day6/input")
	if err != nil {
		panic(err)
	}

	// part1
	cgroups := puzzlein.Clean(groups)
	//findMatchAny(cgroups)

	// part2
	findMatchEvery(cgroups)
}
