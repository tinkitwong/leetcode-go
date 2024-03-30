package merge_two_sorted_lists_21

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{
		Val: 0,
	}
	curr := dummy

	// while either list1 or list2 exists
	for list1 != nil && list2 != nil {
		fmt.Println(list1.Val, list2.Val)
		// if list1 value >= list2 value:
		// - add list2 value to dummy.Next
		// - move to next list2 value
		// - move dummy to next dummy
		if list1.Val >= list2.Val {
			curr.Next = list2
			curr = curr.Next
			list2 = list2.Next
		} else {
			// else (list1 value < list2 value):
			// - add list1 value to dummy.Next
			// - move to next list1 value
			// - move dummy to next dummy
			curr.Next = list1
			curr = curr.Next
			list1 = list1.Next
		}
	}

	// add back remaining nodes
	if list2 != nil {
		curr.Next = list2
	}

	if list1 != nil {
		curr.Next = list1
	}

	return dummy.Next
}
