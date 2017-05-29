package add_two_numbers

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode  {
	var head, tail *ListNode
	carry := 0;
	sum := 0;
	start := true
	for l1 != nil && l2 != nil {
    sum,carry = calculateSumAndCarry(l1, l2, sum, carry)
		if(start) {
			head = &ListNode{Val: sum}
			tail = head
			start = false
		} else {
			tail.Next = &ListNode{Val: sum}
			if(head.Next == nil) {
				head.Next = tail
			}
			tail = tail.Next
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	rem := remaining(l1, l2)

	if(rem == nil) {
		if(carry > 0) {
			tail.Next = &ListNode{Val: carry}
		}
		return head
	}
	tail.Next = addTwoNumbers(rem, &ListNode{Val: carry})
	return head
}
func calculateSumAndCarry(l1 *ListNode, l2 *ListNode, sum int, carry int) (int, int) {
	res := l1.Val + l2.Val + carry
	carry = res / 10
	return res % 10, carry
}

func remaining(l1 *ListNode, l2 *ListNode) *ListNode {
	if (l1 == nil && l2 == nil) {
		return nil;
	} else if (l1 == nil) {
		return l2
	} else {
		return l1
	}
}
