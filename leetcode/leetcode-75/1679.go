package main

import (
	"fmt"
	"sort"
)

/*
Utiliser 2 pointeurs avec la list sorted

*/

func maxOperations(nums []int, k int) int {

	total := 0

	sort.Ints(nums)
	j := len(nums) - 1
	i := 0
	for i < j {
		s := nums[i] + nums[j]
		if s == k {
			total += 1
			i += 1
			j -= 1
			continue
		} else if s < k {
			i += 1
		} else if s > k {
			j -= 1
		}

	}
	return total

}

func main() {
	fmt.Println(maxOperations([]int{3, 1, 3, 4, 3}, 6))
}
