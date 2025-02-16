package btree

import "fmt"

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return findSums(root, 0, 0, targetSum)
}

func findSums(root *TreeNode, currentSum int, sumCount int, targetSum int) int {
	if root == nil {
		fmt.Println("root is nil sumCount =", sumCount)
		return sumCount
	}
	currentSum += root.Val
	fmt.Println("currentSum", currentSum, "root.Val", root.Val, "sumCount", sumCount)
	if currentSum == targetSum {
		sumCount++
		currentSum = 0
	}
	if currentSum > targetSum {
		currentSum = 0
	}

	sumCount = findSums(root.Left, currentSum, sumCount, targetSum)
	sumCount = findSums(root.Right, currentSum, sumCount, targetSum)
	return sumCount
}
