package average_of_levels_in_binary_tree_637

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	// base case
	if root == nil {
		return []float64{0.0}
	}

	queue := []*TreeNode{}
	result := []float64{}

	// add first node into queue
	queue = append(queue, root)

	for len(queue) > 0 {
		sum := 0
		count := 0
		tmp := []*TreeNode{}

		// calculate the current level's sum, and get no. of nodes
		// before pushing in children of this level
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			count += 1
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		queue = tmp
		result = append(result, float64(sum)/float64(count))
	}

	return result
}
