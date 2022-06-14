package _6_reverse_list_node_print

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

func reversePrint(head *ListNode) []int {
	tmp := []int{}
	res := []int{}
	if head == nil {
		return nil
	}
	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}
	for i := len(tmp) - 1; i >= 0; i-- {
		res = append(res, tmp[len(tmp)-1])
		tmp = tmp[:len(tmp)-1]
	}
	return res
}
