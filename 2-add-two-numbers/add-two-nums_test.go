package add_two_numbers

import (
	"testing"
	"fmt"
)

func TestAddTwoNumbers(t *testing.T)  {
	l1 := new(ListNode)
	l1.Val = 5
	l1.Next = &ListNode{Val: 1}
	l1.Next.Next = &ListNode{Val: 2}
	var l2 *ListNode
	l2.Val = 6
	l2.Next = &ListNode{Val: 2}
	l2.Next.Next = &ListNode{Val: 6}
	result := addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Printf("%d", result.Val)
		result = result.Next
	}
}

