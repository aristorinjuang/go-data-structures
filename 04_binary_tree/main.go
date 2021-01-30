package main

import "fmt"

type node struct {
	value byte
	left  *node
	right *node
}

func (n *node) print() {
	fmt.Println(n.value)

	if n.left != nil {
		n.left.print()
	}

	if n.right != nil {
		n.right.print()
	}
}

func (n *node) push(value byte) {
	switch {
	case value < n.value:
		if n.left == nil {
			n.left = &node{value, nil, nil}
		} else {
			n.left.push(value)
		}
	default:
		if n.right == nil {
			n.right = &node{value, nil, nil}
		} else {
			n.right.push(value)
		}
	}
}

func main() {
	tree := &node{5, nil, nil}
	tree.push(3)
	tree.push(7)
	tree.push(1)
	tree.push(2)
	tree.push(6)
	tree.push(8)
	tree.print()
}
