package aoc

import (
	"2023/pkg/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/zyedidia/generic/stack"
)

func Cards() {

	scanner := bufio.NewScanner(utils.ReadInput("d4.input"))

	/* Part2

	For each winning number, you win the next card in fact, et ainsi de suite

	Example

		Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
		4 matching number = cards 2,3,4,5

	We just need to parse it once, record the count of wins per row (card)

	Then for each win, recursively inspect the subsequent rows / memoize ?
	and keep track of how many cards where visited


	 map[
	 	1:4
			2:2
				3:2
					4:1
						5:0
					5:0
				4:1
					5:0
			3:2
				4:1
					5:0
				5:0
			4:1
				5:0
			5:0
		2:2
			3:2
				4:1
					5:0
				5:0
			4:1
				5:0
		3:2
			4:1
				5:0
			5:0
		4:1
			5:0
		5:0
		6:0
	]

	*/
	// get base play
	plays := part2(scanner)
	fmt.Println(plays)
	visit(plays)

}

//func visit(root map[int]int) {
//	visited :=
//	for k,v := range root {
//
//
//	}
//
//}

// A play characterize a play in the games
type play struct {
	index int
	wins  int
	//level int // record the depth in the graph ?
}

func visit(plays map[int]int) {

	s := stack.New[play]()

	// first add all the base cards
	for k, v := range plays {
		p := play{
			index: k,
			wins:  v,
		}
		s.Push(p)
	}

	// track how many where visited
	cards := 0

	// until the stack is empty
	for s.Size() > 0 {
		var toVisit play
		toVisit = s.Pop()
		cards += 1

		// for each play, we need to visit the n "wins" followings play
		for i := toVisit.index + 1; i <= toVisit.wins+toVisit.index; i++ {
			// build the next play to visit
			wins, ok := plays[i]
			if !ok {
				break
			}
			p := play{
				index: i,
				wins:  wins,
			}
			s.Push(p)
		}

	}
	fmt.Println(cards)
}

// parse input and build our base datastructure (root nodes)
func part2(scanner *bufio.Scanner) map[int]int {
	m := map[int]int{}
	i := 1
	for scanner.Scan() {
		line := scanner.Text()

		// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
		left := strings.Split(line, "|")[0]
		right := strings.Split(line, "|")[1]

		// record winning number by index
		win := make([]bool, 100)
		// 41 48 83 86 17
		left = strings.Split(left, ":")[1]

		// parse and sanitize wins
		for _, num := range strings.Split(left, " ") {

			num = strings.Trim(num, " ")
			if num == "" {
				continue
			}
			i, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}

			// store it
			if i > len(win)-1 {
				win = win[:i-len(win)]
			}
			win[i] = true
		}

		// now check the numbers "tirée au sort"
		// parse and sanitize
		t := 0
		for _, num := range strings.Split(right, " ") {
			num = strings.Trim(num, " ")
			if num == "" {
				continue
			}
			i, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}

			if win[i] {
				t += 1
			}
		}
		m[i] = t
		i += 1
	}
	return m
}

func part1(scanner *bufio.Scanner) {
	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
		left := strings.Split(line, "|")[0]
		right := strings.Split(line, "|")[1]

		// record winning number by index
		win := make([]bool, 100)
		// 41 48 83 86 17
		left = strings.Split(left, ":")[1]

		// parse and sanitize wins
		for _, num := range strings.Split(left, " ") {

			num = strings.Trim(num, " ")
			if num == "" {
				continue
			}
			i, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}

			// store it
			if i > len(win)-1 {
				win = win[:i-len(win)]
			}
			win[i] = true
		}

		// now check the numbers "tirée au sort"
		// parse and sanitize
		t := 0
		for _, num := range strings.Split(right, " ") {
			num = strings.Trim(num, " ")
			if num == "" {
				continue
			}
			i, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}

			if win[i] {
				fmt.Println(i)
				if t == 0 {
					t += 1
				} else {
					t *= 2
				}
			}
		}
		total += t
	}
	fmt.Println(total)
}
