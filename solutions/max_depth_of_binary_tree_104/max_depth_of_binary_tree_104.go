package max_depth_of_binary_tree_104

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	// check for empty tree
	if root == nil {
		return 0
	}
	// compute the depth of each subtree (left and right)
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}
