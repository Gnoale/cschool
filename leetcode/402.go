package main

import (
	"fmt"
	"strconv"
)

func removeKdigits(num string, k int) string {
	/*
	   approche avec la monotonic increasing stack ?
	   "1432219"
	   k = 3
	   i  stack k
	   0   1    3
	   1   1    2
	   2   1    1
	   3   1    0
	   4   12
	       122
	       1219

	   "10200"
	   k = 1
	   i  stack k
	   0   1    1
	   1   0    0
	   2   02   0
	   3   020
	   4   0200

	   "1173"
	   k = 2
	   i  stack k
	   0   1    2
	   1   11   2
	   2   11   1
	   3   11   0

	*/

	stack := []int{}

	for i := range num {
		n, _ := strconv.Atoi(string(num[i]))
		if len(stack) == 0 {
			stack = append(stack, n)
		} else if k > 0 {
			top := stack[len(stack)-1]
			if top > n {
				// pop
				stack = stack[:len(stack)-1]
				// and insert
				stack = append(stack, n)
				k--
			} else if top < n {
				// skip and decrease k
				k--
			} else {
				stack = append(stack, n)
			}
		} else {
			stack = append(stack, n)
		}
	}

	fmt.Println(stack)
	result := 0
	for i := len(stack) - 1; i >= 0; i-- {
		result += stack[i] * pow(10, len(stack)-i-1)
	}
	return strconv.Itoa(result)
}

func pow(n, exp int) int {
	if exp == 0 {
		return 1
	}
	base := n
	for i := 1; i < exp; i++ {
		n *= base
	}
	return n
}

func main() {
	fmt.Println(pow(2, 4))

	fmt.Println(removeKdigits("1432219", 3))
	fmt.Println(removeKdigits("10200", 1))
	fmt.Println(removeKdigits("10", 2))
	fmt.Println(removeKdigits("10", 2))
}
