package main

import "fmt"

type node struct {
	value int
	next  *node
}

func reverse(n *node) *node {
	next := n.next
	n.next = nil

	for next != nil {
		tmp := next.next
		next.next = n
		n = next
		next = tmp
	}

	return n
}

func main() {
	root := &node{value: 1}
	root = reverse(root)

	fmt.Println(root)
}
