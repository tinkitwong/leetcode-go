package minimum_absolute_difference_in_bst_530

import "fmt"

func Run() {
	root := &TreeNode{Val: 4}

	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}

	res := getMinimumDifference(root)
	fmt.Println(res)
}
