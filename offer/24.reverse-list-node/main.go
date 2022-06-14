package _24_reverse_list_node

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var fast *ListNode
	slow := head
	for slow != nil {
		next := slow.Next
		slow.Next = fast
		fast = slow
		slow = next
	}
	return fast
}
