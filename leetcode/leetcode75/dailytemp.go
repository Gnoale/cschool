package leetcode75

import (
	"fmt"
)

func dailyTemperatures(temperatures []int) []int {
	response := make([]int, len(temperatures))

	// maintain a monotonic stack in ascending order
	// the first element is always the bigger
	// for each index i
	// if the current value at i is bigger than the latest in the stack,
	// we pop elements from the stack until we find one bigger
	// until a bigger is foudn it means we have the n previous element < currrent
	// we just need to store r[j] = i-j which is the distance for every elements in the stack !
	stack := []int{}

	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && (temperatures[stack[len(stack)-1]] < temperatures[i]) {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Println(stack, i)
			response[j] = i - j // distance between current index and (previous < current)
		}
		stack = append(stack, i)
	}

	return response
}
