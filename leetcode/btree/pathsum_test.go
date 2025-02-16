package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pathSumCases = []struct {
	root      *TreeNode
	targetSum int
	expected  int
}{
	// btree:
	// 10,5,-3,3,2,null,11,3,-2,null,1
	{root: &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: -2,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
		Right: &TreeNode{
			Val: -3,
			Right: &TreeNode{
				Val: 11,
			},
		},
	},
		targetSum: 8,
		expected:  3,
	},
	//[5,4,8,11,null,13,4,7,2,null,null,5,1]
	{root: &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	},
		targetSum: 22,
		expected:  3,
	},
}

func TestPathSum(t *testing.T) {
	for _, testCase := range pathSumCases {
		assert.Equal(t, testCase.expected, pathSum(testCase.root, testCase.targetSum))
	}
}
