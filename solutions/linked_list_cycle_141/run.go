package linked_list_cycle_141

import "fmt"

func Run() {
	values := []int{1}
	head := &ListNode{
		Val: values[0],
	}
	current := head
	for i := 1; i < len(values); i++ {
		newNode := &ListNode{
			Val: values[i],
		}
		current.Next = newNode
		current = current.Next
	}

	// current.Next = head

	res := hasCycleII(head)
	fmt.Println(res)
}
