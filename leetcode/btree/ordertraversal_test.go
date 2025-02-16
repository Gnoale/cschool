package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var levelOrderCases = []struct {
	root     *TreeNode
	expected [][]int
}{
	// [3,9,20,null,null,15,7]
	{root: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, expected: [][]int{{3}, {9, 20}, {15, 7}}},
	// [1]
	{root: &TreeNode{Val: 1}, expected: [][]int{{1}}},
	// []
	{root: nil, expected: [][]int{}},
	// [1,2,3,4,null,null,5]
	{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 5}}}, expected: [][]int{{1}, {2, 3}, {4, 5}}},
}

func TestLevelOrder(t *testing.T) {
	for _, testCase := range levelOrderCases {
		assert.Equal(t, testCase.expected, levelOrder(testCase.root))
	}
}
