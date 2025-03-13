package singly

import "fmt"

type node struct {
	next *node
	val  int
}

func LinkedSinglyOps() {
	head := &node{val: 1}
	head.next = &node{val: 2}
	head.next.next = &node{val: 3}
	head.next.next.next = &node{val: 4}
	head.next.next.next.next = &node{val: 5}

	printList(head)

}

func printList(head *node) {
	curr := head

	for curr != nil {
		if curr.next == nil {
			fmt.Print(curr.val)
		} else {
			fmt.Print(curr.val, " -> ")
		}
		curr = curr.next
	}

}
