package main

import (
	"fmt"
	"math"
)

type node struct {
	value byte
	next  *node
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

func (n *node) print() {
	current := n

	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}

	fmt.Println()
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
			current.value = current.next.value
			current = current.next
		}

		previous.next = nil
	}
}

func (n *node) front() byte {
	return n.value
}

func (n *node) back() byte {
	current := n

	for current.next != nil {
		current = current.next
	}

	return current.value
}

func main() {
	queue := node{1, nil}
	queue.push(2)
	queue.push(3)
	queue.pop()
	queue.print()
	fmt.Println(queue.back())
	fmt.Println(queue.front())
	fmt.Println(queue.size())
	fmt.Println(queue.empty())
}
