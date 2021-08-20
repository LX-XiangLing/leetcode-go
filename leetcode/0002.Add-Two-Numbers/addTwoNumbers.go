package _002_Add_Two_Numbers

/*
	题目链接:https://leetcode-cn.com/problems/add-two-numbers/
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l103 := &ListNode{Val: 3}
	l102 := &ListNode{Val: 4, Next: l103}
	l101 := &ListNode{Val: 2, Next: l102}
	l203 := &ListNode{Val: 8}
	l202 := &ListNode{Val: 0, Next: l203}
	l201 := &ListNode{Val: 7, Next: l202}
	addTwoNumbers(l101, l201)
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var tail *ListNode
	var head *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}
