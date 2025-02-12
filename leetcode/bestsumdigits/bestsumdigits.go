package bestsumdigits

import "sort"

func sumOfDigits(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

func MaximumSum(nums []int) int {

	//sort.Ints(nums)
	// [7,13,18,36,43]
	// [2,2,5,9]
	max := 0

	sort.Slice(nums, func(i, j int) bool {
		if sumOfDigits(nums[i]) < sumOfDigits(nums[j]) {
			return true
		}
		if sumOfDigits(nums[i]) == sumOfDigits(nums[j]) {
			return nums[i] < nums[j]
		}
		return false
	})
	j := 1
	for i := 0; i < len(nums)-1 && j < len(nums); i++ {
		if sumOfDigits(nums[i]) == sumOfDigits(nums[j]) {
			sum := nums[i] + nums[j]
			if sum > max {
				max = sum
			}
		}
		j++
	}
	if max == 0 {
		return -1
	}
	return max
}
