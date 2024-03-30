package merge_two_sorted_lists_21

import "fmt"

func Run() {
	list1_values := []int{1, 2, 4}
	list2_values := []int{1, 3, 4}

	list1_head := &ListNode{
		Val: list1_values[0],
	}
	list2_head := &ListNode{
		Val: list1_values[0],
	}
	list1_current := list1_head
	list2_current := list2_head

	// create list1
	for i := 1; i < len(list1_values); i++ {
		newNode := &ListNode{
			Val: list1_values[i],
		}
		list1_current.Next = newNode
		list1_current = list1_current.Next
	}

	// create list2
	for i := 1; i < len(list2_values); i++ {
		newNode := &ListNode{
			Val: list2_values[i],
		}
		list2_current.Next = newNode
		list2_current = list2_current.Next
	}

	res := mergeTwoLists(list1_head, list2_head)
	fmt.Println(res)
}
