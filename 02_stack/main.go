package main

import (
	"fmt"
	"math"
)

type node struct {
	value byte
	next  *node
}

func (n *node) print() {
	current := n

	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}

	fmt.Println()
}

func (n *node) size() uint32 {
	current := n
	var size uint32 = 0

	for current != nil {
		current = current.next
		size++
	}

	return size
}

func (n *node) empty() bool {
	return n.size() == 0
}

func (n *node) push(value byte) {
	if n.size() < math.MaxInt32 {
		current := n

		for current.next != nil {
			current = current.next
		}

		current.next = &node{value, nil}
	}
}

func (n *node) pop() {
	if !n.empty() {
		var previous, current *node = nil, n

		for current.next != nil {
			previous = current
			current = current.next
		}

		previous.next = nil
	}
}

func (n *node) top() byte {
	current := n

	for current.next != nil {
		current = current.next
	}

	return current.value
}

func main() {
	stack := node{1, nil}
	stack.push(2)
	stack.push(3)
	stack.pop()
	stack.print()
	fmt.Println(stack.top())
	fmt.Println(stack.size())
	fmt.Println(stack.empty())
}
