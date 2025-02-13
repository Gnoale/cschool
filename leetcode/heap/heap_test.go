package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapAdd(t *testing.T) {
	heap := NewHeap(10)
	heap.Add(5)
	heap.Add(3)
	heap.Add(8)
	heap.Add(1)
	heap.Add(10)
	heap.Add(-1)
	heap.Add(9)

	assert.Equal(t, -1, heap.Nodes[0].Value)
	assert.Equal(t, 3, heap.Nodes[1].Value)
	assert.Equal(t, 1, heap.Nodes[2].Value)
	assert.Equal(t, 5, heap.Nodes[3].Value)
	assert.Equal(t, 10, heap.Nodes[4].Value)
	assert.Equal(t, 8, heap.Nodes[5].Value)
	assert.Equal(t, 9, heap.Nodes[6].Value)
}

var popTestCases = []struct {
	nums     []int
	expected []int
}{
	{[]int{5, 3, 8, 1, 10, -1, 9}, []int{-1, 1, 3, 5, 8, 9, 10}},
	{[]int{7, 11, 10, 3, 11}, []int{3, 7, 10, 11, 11}},
	{[]int{84, 74, 86}, []int{74, 84, 86}},
}

func TestHeapPop(t *testing.T) {
	for _, testCase := range popTestCases {
		heap := NewHeap(len(testCase.nums))
		for _, v := range testCase.nums {
			heap.Add(v)
		}
		for _, v := range testCase.expected {
			assert.Equal(t, v, heap.Pop())
		}
	}
}
