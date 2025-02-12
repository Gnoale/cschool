package bestsumdigits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	nums     []int
	expected int
}{
	{[]int{18, 43, 36, 13, 7}, 54},
	{[]int{10, 12, 19, 14}, -1},
	{[]int{9, 2, 2, 5}, 4},
	{[]int{368, 369, 307, 304, 384, 138, 90, 279, 35, 396, 114, 328, 251, 364, 300, 191, 438, 467, 183}, 835},
}

func TestMaximumSum(t *testing.T) {
	for _, testCase := range testCases {
		assert.Equal(t, testCase.expected, MaximumSum(testCase.nums))
	}
}
