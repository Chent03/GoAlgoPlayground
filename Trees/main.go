package main

import (
	"fmt"
)

type Node struct {
	left  *Node
	right *Node
	data  int
}

func (N *Node) Insert(value int) {
	if value <= N.data {
		if N.left == nil {
			N.left = &Node{
				data: value,
			}
		} else {
			N.left.Insert(value)
		}
	} else {
		if N.right == nil {
			N.right = &Node{
				data: value,
			}
		} else {
			N.right.Insert(value)
		}
	}
}

func (N *Node) Contains(value int) bool {
	if value == N.data {
		return true
	} else if value < N.data {
		if N.left == nil {
			return false
		} else {
			return N.left.Contains(value)
		}
	} else {
		if N.right == nil {
			return false
		} else {
			return N.right.Contains(value)
		}
	}
}

func (N *Node) PrintInOrder() {
	if N.left != nil {
		N.left.PrintInOrder()
	}
	fmt.Println(N.data)
	if N.right != nil {
		N.right.PrintInOrder()
	}
}

func main() {
	tree := Node{
		data: 10,
	}

	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(8)

	tree.PrintInOrder()
}
