package main

import (
	"fmt"
)

type Node struct {
	prev *Node
	next *Node
	data interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Append(data interface{}) {

	if L.head == nil {
		L.head = &Node{
			data: data,
		}
		return
	}
	l := L.head
	for l.next != nil {
		l = l.next
	}

	l.next = &Node{
		data: data,
	}
}

func (L *List) Prepend(data interface{}) {
	list := &Node{
		data: data,
	}
	list.next = L.head
	L.head = list
}

func (L *List) Display() {
	list := L.head
	for list != nil {
		fmt.Printf("%v ->", list.data)
		list = list.next
	}
}

func (L *List) Delete(data interface{}) {
	if L.head == nil {
		return
	}
	if L.head.data == data {
		L.head = L.head.next
		return
	}

	l := L.head
	for l.next != nil {
		if l.next.data == data {
			l.next = l.next.next
			return
		}
		l = l.next
	}
}

func main() {
	link := List{}
	link.Append(1)
	link.Append(2)
	link.Display()
	fmt.Println("\n==============================")
	link.Prepend(3)
	link.Display()
	fmt.Println("\n==============================")
	link.Delete(1)
	link.Display()
	fmt.Println("\n==============================")
	link.Delete(3)
	link.Display()
}
