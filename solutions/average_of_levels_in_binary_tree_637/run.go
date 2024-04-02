package average_of_levels_in_binary_tree_637

import "fmt"

func Run() {
	root := &TreeNode{Val: 3}

	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	res := averageOfLevels(root)
	fmt.Println(res)
}
