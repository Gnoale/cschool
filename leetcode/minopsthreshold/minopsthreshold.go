package minopsthreshold

import (
	"leetcode/heap"
)

func minOperations(nums []int, k int) int {

	heap := heap.NewHeap(len(nums))
	for _, v := range nums {
		heap.Add(v)
	}

	operationsCount := 0
	if min := heap.Peek(); min >= k {
		return operationsCount
	}

	for {
		if heap.LastIndex == 0 {
			break
		}
		x, y := heap.Pop(), heap.Pop()
		min, max := x, y
		sum := min*2 + max

		heap.Add(sum)
		operationsCount++

		if min := heap.Peek(); min >= k {
			break
		}
	}
	return operationsCount
}
