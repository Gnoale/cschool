package leetcode75

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	origin := head
	if head == nil {
		return nil
	}
	n := 0

	// first count the nodes
	for head != nil {
		n += 1
		head = head.Next
	}

	var ancestor *ListNode = nil
	head = origin
	i := -1
	for head != nil {
		i += 1
		if i == n/2 {
			if ancestor == nil {
				return nil
			}
			if head.Next != nil {
				ancestor.Next = head.Next
			} else {
				ancestor.Next = nil
			}
			break
		}
		ancestor = head
		head = head.Next
	}
	return origin

}

func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	var ancestor *ListNode = nil
	var child *ListNode = nil
	for head.Next != nil {
		child = head.Next
		head.Next = ancestor
		ancestor = head
		head = child
	}
	head.Next = ancestor
	return head

}
