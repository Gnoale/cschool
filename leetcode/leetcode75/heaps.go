package leetcode75

import "container/heap"

type IntMaxHeap []int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMaxHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() any {
	size := len(*h)
	latest := (*h)[size-1]
	*h = (*h)[:size-1]
	return latest
}

func findKthLargest(nums []int, k int) int {

	intMaxHeap := &IntMaxHeap{}
	heap.Init(intMaxHeap)

	for _, num := range nums {
		heap.Push(intMaxHeap, num)
	}

	for i := 0; i < k-1; i++ {
		heap.Pop(intMaxHeap)
	}

	return heap.Pop(intMaxHeap).(int)
}
