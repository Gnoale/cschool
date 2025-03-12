package btree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return findMaxDepth(root, 0)
}

func findMaxDepth(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	leftDepth := findMaxDepth(root.Left, depth+1)
	rightDepth := findMaxDepth(root.Right, depth+1)
	return max(leftDepth, rightDepth)
}

/*
represent [3,1,4,3,null,1,5] as btree
given a btree, a good node is a node whose value is greater than or equal to the maximum value along the path from its ancestor.
exemple:

		3
	   / \
	  1   4
	 /   /  \
	3   1    5
*/
func findGoodNodes(node *TreeNode, max int, count int) int {
	if node == nil {
		return count
	}
	if node.Val >= max {
		max = node.Val
		count++
	}
	count = findGoodNodes(node.Left, max, count)
	count = findGoodNodes(node.Right, max, count)
	return count
}

/*
given a binary tree, a zigzag path is a path that starts at any node and alternates between left and right children.
exemple:

		1
	   / \
	  1   1
	 / \
	1   1
	   /
	  1

	longest zigzag path is 3
*/
func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}
}
