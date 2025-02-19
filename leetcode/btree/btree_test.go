package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var maxDepthCases = []struct {
	root     *TreeNode
	expected int
}{
	{
		root:     &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
		expected: 3,
	},
	{
		root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
		expected: 2,
	},
	{
		root:     nil,
		expected: 0,
	},
}

func TestMaxDepth(t *testing.T) {
	for _, testCase := range maxDepthCases {
		assert.Equal(t, testCase.expected, maxDepth(testCase.root))
	}
}

var levelOrderCases = []struct {
	root     *TreeNode
	expected [][]int
}{
	{root: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, expected: [][]int{{3}, {9, 20}, {15, 7}}},
	{root: &TreeNode{Val: 1}, expected: [][]int{{1}}},
	{root: nil, expected: [][]int{}},
	{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 5}}}, expected: [][]int{{1}, {2, 3}, {4, 5}}},
}

func TestLevelOrder(t *testing.T) {
	for _, testCase := range levelOrderCases {
		assert.Equal(t, testCase.expected, levelOrder(testCase.root))
	}
}
