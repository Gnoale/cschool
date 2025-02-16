package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	root     *TreeNode
	expected int
}{
	{
		root:     &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
		expected: 3,
	},
}

func TestMaxDepth(t *testing.T) {
	for _, testCase := range testCases {
		assert.Equal(t, testCase.expected, maxDepth(testCase.root))
	}
}
