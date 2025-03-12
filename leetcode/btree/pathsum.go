package btree

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	prefixSums := map[int]int{0: 1}
	return findSums(root, 0, 0, targetSum, prefixSums)
}

func findSums(node *TreeNode, sum int, paths int, k int, prefixSums map[int]int) int {
	sum += node.Val
	// along the current path there are 2 possibilities satisfying the target sum:
	// the current node value sum to the target sum, in this case the differennce is 0 and this is a match
	// the current sum - target equal some sum along the path, it means that we found a subtree satisfying the target sum
	if count, ok := prefixSums[sum-k]; ok {
		paths += count
	}
	prefixSums[sum]++

	if node.Left != nil {
		paths = findSums(node.Left, sum, paths, k, prefixSums)
	}
	if node.Right != nil {
		paths = findSums(node.Right, sum, paths, k, prefixSums)
	}

	// prune the current leaf node sum
	prefixSums[sum]--
	return paths
}

/*
This technique allow to explore all paths in the tree, maintain a count for each level
And recursvely prune each node to avoid counting the same path multiple times

*/
