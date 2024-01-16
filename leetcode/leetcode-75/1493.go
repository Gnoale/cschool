package main

import "fmt"

/*
Given a binary array nums, you should delete one element from it.

Return the size of the longest non-empty subarray containing only 1's in the resulting array. Return 0 if there is no such subarray.

Example 2:

Input: nums = [0,1,1,1,0,1,1,0,1]
Output: 5
Explanation: After deleting the number in position 4, [0,1,1,1,1,1,0,1] longest subarray with value of 1's is [1,1,1,1,1].
*/

/*
Ressemble à 1004. Max Consecutive Ones III

On commence avec left, right = -1

On peut flip 1 seul 0 dans la window




*/

func longestSubarray(nums []int) int {
	left, right := -1, -1
	k := 1
	L := 0

	for right < len(nums)-1 {
		// move window 1 step right
		right++
		if nums[right] == 0 {
			k--
		}
		// if more than a zero in the window
		// slide left until it is good
		for k < 0 {
			left++
			if nums[left] == 0 {
				k++
			}
		}
		if k >= 0 {
			if right-left > L {
				//fmt.Println(left, right)
				L = right - left - 1
			}
		}

	}

	return L

}

func main() {
	fmt.Println(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}))
	fmt.Println(longestSubarray([]int{1, 1, 0, 1}))
	fmt.Println(longestSubarray([]int{1, 1, 1}))
}
