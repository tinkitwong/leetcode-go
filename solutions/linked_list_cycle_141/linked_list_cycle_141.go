package linked_list_cycle_141

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	// create set of ListNodes that has been seen
	seen := map[*ListNode]struct{}{}

	// iterate through each node
	for head != nil {
		fmt.Println("current value:", head.Val)
		// check if current node has been visited
		if _, ok := seen[head]; ok {
			// return true as this indicates a cycle
			return true
		} else {
			// else add unseen ListNode to set
			seen[head] = struct{}{}
		}
		head = head.Next
	}

	return false
}

// Floyd's Tortoise and Hare algorithm
func hasCycleII(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}
