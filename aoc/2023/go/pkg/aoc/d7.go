package aoc

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

var CardsV = []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}

var Rank = [][]int{{5, 0}, {4, 0}, {3, 2}, {3, 0}, {2, 3}, {2, 0}, {1, 0}}

type hand struct {
	hand  string
	score int
}

func parseD7(scanner *bufio.Scanner) []hand {
	hands := []hand{}
	for scanner.Scan() {
		line := scanner.Text()
		sline := strings.Split(line, " ")
		score, err := strconv.Atoi(sline[1])
		if err != nil {
			panic(err)
		}
		hands = append(hands, hand{
			hand:  sline[0],
			score: score,
		})
	}
	return hands
}

func CamelPoker(scanner *bufio.Scanner) {
	hands := parseD7(scanner)
	/*
		Combinaision are ranked from strongest to weakest
			AAAAA (5)
			AA8AA (4)
			23332 (3,2)
			TTT98 (3,0)
			23432 (2,2)
			A23A4 (2)
			A23B4 (1)

		If ex aequo, compare cards from left to right until one is stronger than the other

	*/
	// sort bid per rank
	sort.SliceStable(hands, func(i, j int) bool {
		// records cards occurence in a map
		cardsI := map[string]int{}
		for k, w := 0, 0; k < len(hands[i].hand); k += w {
			r, width := utf8.DecodeRuneInString(hands[i].hand[k:])
			cardsI[string(r)] += 1
			w = width
		}
		cardsJ := map[string]int{}
		for k, w := 0, 0; k < len(hands[j].hand); k += w {
			r, width := utf8.DecodeRuneInString(hands[j].hand[k:])
			cardsJ[string(r)] += 1
			w = width
		}
		ok := CompareHands(cardsI, cardsJ, hands[i].hand, hands[j].hand)
		//fmt.Println(hands[i].hand, "<", hands[j].hand, ok)
		return ok
	})

	fmt.Println(hands)
	t := 0
	for i := 0; i < len(hands); i++ {
		//runtime.Breakpoint()
		t += (i + 1) * hands[i].score
	}
	fmt.Println("score =", t)
}

// count occurence of cards in a hand and return the score
func rankHand(hand map[string]int) int {
	r := 0
	for _, c := range hand {
		if c == 0 || c == 1 {
			continue
		}
		if c == 5 {
			r = 7
			break
		}
		if c == 4 {
			r = 6
			break
		}
		if c == 3 {
			// 2,3
			if r == 2 {
				r = 5
				break
			}
			// 3,0
			r = 4
		}
		if c == 2 {
			// 3,2
			if r == 4 {
				r = 5
				break
			}
			// 2,2
			if r == 2 {
				r = 3
				break
			}
			// 2,0
			r = 2
		}
	}
	return r
}

func joker(r int, J int) int {
	// use the Joker in such cases
	switch r {
	case 6, 5: // 4,1 3,2
		return 7 // 5
	case 4: // 3,1,1
		return 6 // 4,1
	case 3: // 2,2,1
		if J == 2 {
			return 6 // 4,1
		}
		if J == 1 {
			return 5 // 3,2
		}
	case 2: // 2,1,1,1
		return 4 // 3,1,1 > 2,2,1
	case 0:
		return 2 // 2,1,1,1
	}
	return r
}

// Compare 2 hands relatively given their set of cards
func CompareHands(left, right map[string]int, hleft, hright string) bool {

	rankL := rankHand(left)
	rankR := rankHand(right)

	// part2 Jokers
	JL := left["J"]
	if JL > 0 {
		rankL = joker(rankL, JL)
	}
	JR := right["J"]
	if JR > 0 {
		rankR = joker(rankR, JR)
	}

	if rankL != rankR {
		return rankL < rankR
	}
	// compare cards 1 by 1 until one wins
	for i := 0; i < len(hleft); i++ {

		if hleft[i] != hright[i] {
			s1, s2 := 0, 0
			for j, card := range CardsV {
				if card == string(hleft[i]) {
					s1 = j
				}
				if card == string(hright[i]) {
					s2 = j
				}
			}
			return s1 < s2
		}
	}
	return false
}
