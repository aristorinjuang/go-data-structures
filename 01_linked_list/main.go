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
			current = current.next
		}

		previous.next = nil
	}
}

func (n *node) back() byte {
	current := n

	for current.next != nil {
		current = current.next
	}

	return current.value
}

func (n *node) front() byte {
	return n.value
}

func (n *node) at(index uint32) byte {
	current := n
	var i uint32
	var result byte

	for current.next != nil && i < index {
		current = current.next
		i++
	}

	result = current.value

	if index > n.size() {
		result = 0
	}

	return result
}

func (n *node) insert(value byte, index uint32) {
	if n.size() < math.MaxInt32 {
		switch {
		case index <= 0, index > n.size():
			n.push(value)
		default:
			current := n
			var i uint32

			for current.next != nil && i < index-1 {
				current = current.next
			}

			tmp := &node{current.value, nil}
			next := current.next

			for next != nil {
				tmp.next = &node{next.value, nil}
				next = next.next
			}

			current.value = value
			current.next = tmp
		}
	}
}

func (n *node) erase(index uint32) {
	if !n.empty() && index <= n.size() {
		switch {
		case index >= n.size():
			var previous, current *node = nil, n

			for current.next != nil {
				previous = current
				current = current.next
			}

			previous.next = nil
		default:
			var previous, current *node = nil, n
			var i uint32

			for current.next != nil && i < index {
				previous = current
				current = current.next
				i++
			}

			previous.next = current.next
		}
	}
}

func main() {
	list := node{1, nil}
	list.push(2)
	list.push(3)
	list.pop()
	list.insert(4, 2)
	list.erase(2)
	list.print()
	fmt.Println(list.back())
	fmt.Println(list.front())
	fmt.Println(list.size())
	fmt.Println(list.empty())
	fmt.Println(list.at(2))
}
