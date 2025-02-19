package leetcode75

import (
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
		} else if s < k {
			i += 1
		} else {
			j -= 1
		}
	}
	return total
}
