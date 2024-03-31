package same_tree_100

import "fmt"

func buildTree(vals []int, index int) *TreeNode {
	if index >= len(vals) || vals[index] == -1 {
		return nil
	}

	node := &TreeNode{
		Val: vals[index],
	}
	node.Left = buildTree(vals, 2*index+1)
	node.Right = buildTree(vals, 2*index+2)
	return node
}

func Run() {
	p := &TreeNode{
		Val: 1,
	}

	q := &TreeNode{
		Val: 1,
	}
	q.Right = &TreeNode{
		Val: 2,
	}

	res := isSameTree(p, q)
	fmt.Println(res)
}
