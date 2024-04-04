package minimum_absolute_difference_in_bst_530

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	prev := (*TreeNode)(nil)
	res := math.Inf(1)

	dfs(root, &prev, &res)

	return int(res)
}

// They need to be passed as pointers because you want to modify
// their values inside the dfs function and have those modifications
// reflected in the calling function getMinimumDifference.
func dfs(node *TreeNode, prev **TreeNode, res *float64) {
	if node == nil {
		return
	}

	dfs(node.Left, prev, res)

	// process current node -- minimise the difference
	if *prev != nil {
		*res = math.Min(*res, float64(node.Val-(*prev).Val))
	}

	// set prev node to current node
	*prev = node
	dfs(node.Right, prev, res)
}
