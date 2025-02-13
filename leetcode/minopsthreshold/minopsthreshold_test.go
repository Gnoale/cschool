package minopsthreshold

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	nums     []int
	k        int
	expected int
}{
	{[]int{2, 11, 10, 1, 3}, 10, 2},
	{[]int{1, 1, 2, 4, 9}, 20, 4},
	{[]int{42, 46}, 42, 0},
	{[]int{74, 22, 40, 86}, 86, 2},
}

func TestMinOperations(t *testing.T) {
	for _, testCase := range testCases {
		fmt.Println("testCase", testCase)
		assert.Equal(t, testCase.expected, minOperations(testCase.nums, testCase.k))
	}
}
