package linkedlist

import "fmt"

// Append, Prepend, InsertAfter(node, value), Delete(value), Find(value).

type Node[T any] struct {
	next *Node[T]
	prev *Node[T]
	val  T
}

type LinkedList[T any] struct {
	head   *Node[T]
	length int
}

func (list *LinkedList[T]) Append(newVal T) {
	newNode := &Node[T]{nil, nil, newVal}
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

func (list LinkedList[T]) Print() {
	fmt.Print("head->")
	cur := list.head
	for cur != nil {
		fmt.Printf("[%v]->", cur.val)
		cur = cur.next
	}
	fmt.Println("end")
}
