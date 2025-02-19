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
