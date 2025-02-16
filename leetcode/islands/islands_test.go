package islands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	grid     [][]byte
	expected int
}{
	{
		grid:     [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}},
		expected: 1,
	},
	{
		grid:     [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}},
		expected: 3,
	},
}

func TestNumIslands(t *testing.T) {
	for _, testCase := range testCases {
		result := numIslands(testCase.grid)
		assert.Equal(t, testCase.expected, result)
	}
}
