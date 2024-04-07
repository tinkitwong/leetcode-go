package convert_sorted_array_to_bst_108

import "fmt"

func Run() {
	nums := []int{-10, -3, 0, 5, 9}

	res := sortedArrayToBST(nums)

	// inorder traverse to check bst
	dfs(res)
}

func dfs(node *TreeNode) {
	if node == nil {
		return
	}

	dfs(node.Left)

	fmt.Println(node.Val)

	dfs(node.Right)
}
