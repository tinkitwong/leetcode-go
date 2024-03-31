package max_depth_of_binary_tree_104

import "fmt"

func Run() {
	head := &TreeNode{
		Val: 1,
	}
	head.Left = &TreeNode{
		Val: 2,
	}

	head.Right = &TreeNode{
		Val: 3,
	}

	head.Left.Left = &TreeNode{
		Val: 4,
	}

	head.Left.Right = &TreeNode{
		Val: 5,
	}

	res := maxDepth(head)
	fmt.Println(res)
}
