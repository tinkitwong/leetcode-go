package convert_sorted_array_to_bst_108

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// convert ASC sorted nums into height-balanced bst
func sortedArrayToBST(nums []int) *TreeNode {

	return helper(0, len(nums)-1, nums)
}

func helper(left int, right int, nums []int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper(left, mid-1, nums)
	root.Right = helper(mid+1, right, nums)
	return root
}
