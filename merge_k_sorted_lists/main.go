package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// func mergeKLists(lists []*ListNode) *ListNode {
// 	root := &ListNode{}
// 	curr := root.Next

// 	for _, list := range lists {
// 		l := list
// 		for l != nil {
// 			if curr == root.Next && (curr == nil || l.Val < curr.Val) {
// 				if curr == nil {
// 					curr = &ListNode{Val: l.Val}
// 				} else if l.Val < curr.Val {
// 					temp := &ListNode{Val: l.Val}
// 					temp.Next = curr
// 					curr = temp
// 				}

// 				root.Next = curr
// 				l = l.Next

// 				continue
// 			}

// 			if curr.Next == nil {
// 				curr.Next = &ListNode{Val: l.Val}
// 			} else if l.Val <= curr.Next.Val {
// 				temp := &ListNode{Val: l.Val}
// 				temp.Next = curr.Next
// 				curr.Next = temp
// 			} else {
// 				curr = curr.Next
// 				continue
// 			}

// 			l = l.Next
// 			curr = curr.Next
// 		}

// 		curr = root.Next
// 	}

// 	return root.Next
// }

type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int {
	return len(h)
}

func (h ListNodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h ListNodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ListNodeHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() any {
	hd := *h
	l := len(hd)
	last := hd[l-1]
	*h = hd[:l-1]
	return last
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &ListNodeHeap{}
	heap.Init(h)

	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	root := &ListNode{}
	curr := root
	for h.Len() > 0 {
		l := heap.Pop(h).(*ListNode)
		curr.Next = l
		curr = curr.Next

		if l.Next != nil {
			heap.Push(h, l.Next)
		}
	}

	return root.Next
}

func main() {
	lists := []*ListNode{
		buildList([]int{1}),
		buildList([]int{-1, 0}),
		buildList([]int{-2, 0}),
		// buildList([]int{2, 6}),
	}

	result := mergeKLists(lists)
	printList(result)
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	curr := head
	for _, v := range nums[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}
