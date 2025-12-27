package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	first := &ListNode{}
	last := first

	carry := 0

	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		last.Next = &ListNode{Val: sum}
		last = last.Next

		if sum >= 10 {
			last.Val -= 10
			carry = 1
		} else {
			carry = 0
		}
		// carry = sum / 10
		// last.Next = &ListNode{Val: sum % 10}
		// last = last.Next
	}

	if carry == 1 {
		last.Next = &ListNode{Val: 1}
	}

	return first.Next
}

func print(root *ListNode) {
	if root == nil {
		return
	}

	fmt.Println("Node: ", root.Val)
	print(root.Next)
}

func main() {
	l1 := &ListNode{Val: 9}
	l1.Next = &ListNode{Val: 9}
	l1.Next.Next = &ListNode{Val: 9}
	l1.Next.Next.Next = &ListNode{Val: 9}

	l2 := &ListNode{Val: 9}
	l2.Next = &ListNode{Val: 9}
	l2.Next.Next = &ListNode{Val: 9}

	first := addTwoNumbers(l1, l2)
	print(first)
}
