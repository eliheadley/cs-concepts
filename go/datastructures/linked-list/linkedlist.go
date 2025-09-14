package main

import "fmt"

// Append, Prepend, InsertAfter(node, value), Delete(value), Find(value).

type Node struct {
	next *Node
	val  int
}

type LinkedList struct {
	head   *Node
	length int
}

func (list *LinkedList) add(newVal int) {
	newNode := &Node{nil, newVal}
	if list.head == nil {
		list.head = newNode
		list.length++
		return
	}

	cur := list.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = newNode
	list.length++
}

func (list LinkedList) print() {
	fmt.Print("head->")
	cur := list.head
	for cur != nil {
		fmt.Printf("[%v]->", cur.val)
		cur = cur.next
	}
	fmt.Println("end")
}

func main() {
	list := LinkedList{}
	list.print()

	list.add(1)
	list.add(2)
	list.add(3)
	list.add(4)
	list.add(5)

	list.print()
}
