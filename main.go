package main

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// keep track of total
// keep track of last_node
// if last is > curr
//   subtract from total

func goodNodes(root *TreeNode) int {
	return dfs(root, root.Val)
}

// keep track of the max, root value with v
func dfs(node *TreeNode, v int) int {
	if node == nil {
		return 0
	}
	// keep track of the max
	curMax := max(node.Val, v)
	// increment if the node is good
	if node.Val >= v {
		return 1 + dfs(node.Left, curMax) + dfs(node.Right, curMax)
	}
	// otherwise do not increment if the node is bad
	return dfs(node.Left, curMax) + dfs(node.Right, curMax)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
