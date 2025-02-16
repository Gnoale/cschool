package btree

func levelOrder(root *TreeNode) [][]int {
	// here we use BFS to traverse the tree level by level
	// a queue is used to store the nodes level by level (FIFO)
	queue := []*TreeNode{}
	if root == nil {
		return [][]int{}
	}

	queue = append(queue, root)
	nodesLevel := map[*TreeNode]int{}
	nodesLevel[root] = 0

	// the nodes values at each level needs to be added by order of run through the queue
	// [1,2,3,4,null,null,5]
	//       1
	//      / \
	//     2   3
	//    /     \
	//   4       5
	// Level order traversal: [[1], [2,3], [4,5]]

	order := [][]int{{root.Val}}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			continue
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
			nodesLevel[node.Left] = nodesLevel[node] + 1
			if len(order) <= nodesLevel[node.Left] {
				order = append(order, []int{node.Left.Val})
			} else {
				order[nodesLevel[node.Left]] = append(order[nodesLevel[node.Left]], node.Left.Val)
			}
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			nodesLevel[node.Right] = nodesLevel[node] + 1
			if len(order) <= nodesLevel[node.Right] {
				order = append(order, []int{node.Right.Val})
			} else {
				order[nodesLevel[node.Right]] = append(order[nodesLevel[node.Right]], node.Right.Val)
			}
		}
	}

	return order
}
