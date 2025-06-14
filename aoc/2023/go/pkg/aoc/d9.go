package aoc

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func parseD9(scanner *bufio.Scanner) [][]int64 {
	seq := [][]int64{}
	for scanner.Scan() {
		s := []int64{}
		for _, n := range strings.Split(scanner.Text(), " ") {
			num, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			s = append(s, int64(num))
		}
		seq = append(seq, s)
	}
	return seq
}

func regress(sequence []int64, stack [][]int64) [][]int64 {
	interval := []int64{}
	for i := 1; i < len(sequence); i++ {
		interval = append(interval, sequence[i]-sequence[i-1])
	}
	stack = append(stack, interval)
	sequence = stack[len(stack)-1]
	if interval[len(interval)-1] != 0 {
		return regress(sequence, stack)
	}

	fmt.Println("All Zeros ?", interval)
	return stack
}

func FindSequences(scanner *bufio.Scanner) {
	baseSequences := parseD9(scanner)

	var total int64

	for _, sequence := range baseSequences {
		base := sequence // save
		stacks := regress(sequence, [][]int64{})
		/*
			[
				[-1 -5 -5 2 19 49 95 160 247 359 499 670 875 1117 1399 1724 2095 2515 2987 3514]
				[-4 0 7 17 30 46 65 87 112 140 171 205 242 282 325 371 420 472 527]
				[4 7 10 13 16 19 22 25 28 31 34 37 40 43 46 49 52 55] *
				[3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3]
			]
		*/
		fmt.Println(base)
		for i := len(stacks) - 2; i > 0; i-- { // skip [0,0,...]
			stacks[i-1] = append(stacks[i-1], stacks[i-1][len(stacks[i-1])-1]+stacks[i][len(stacks[i])-1])
		}

		base = append(base, base[len(base)-1]+stacks[0][len(stacks[0])-1])
		latest := base[len(base)-1]
		if latest < 0 {
			latest = -latest
		}
		total += latest
		fmt.Println(base, total)
		fmt.Println()
	}
	fmt.Println(total)
}
