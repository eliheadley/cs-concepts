package linkedlist

import (
	"errors"
	"fmt"
)

// Append, Prepend, InsertAfter(node, value), Delete(value), Find(value).

var IndexOutOfBounds = errors.New("Index is out of bounds for the current size of the list")

type Node[T any] struct {
	next *Node[T]
	prev *Node[T]
	val  T
}

type LinkedList[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func (list *LinkedList[T]) Append(newVal T) {
	newNode := &Node[T]{nil, nil, newVal}
	if list.head == nil && list.tail == nil {
		list.head = newNode
		list.tail = newNode
		list.length++
		return
	}

	list.tail.next = newNode
	newNode.prev = list.tail
	list.tail = newNode
	list.length++
}

func (list *LinkedList[T]) Prepend(newVal T) {
	newNode := &Node[T]{nil, nil, newVal}
	if list.head == nil && list.tail == nil {
		list.head = newNode
		list.tail = newNode
		list.length++
		return
	}

	list.head.prev = newNode
	newNode.next = list.head
	list.head = newNode
	list.length++
}

func (list *LinkedList[T]) InsertAfter(index int, newVal T) error {
	if index < 0 || index >= list.length {
		return IndexOutOfBounds
	}

	// Inserting after the last node → append
	if index == list.length-1 {
		list.Append(newVal)
		return nil
	}

	// Find the node at position index
	cur := list.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	next := cur.next
	newNode := &Node[T]{val: newVal}

	newNode.prev = cur
	newNode.next = next
	cur.next = newNode
	if next != nil {
		next.prev = newNode
	}

	list.length++
	return nil
}

func (list LinkedList[T]) Print() {
	fmt.Print("head->")
	cur := list.head
	for cur != nil {
		fmt.Printf("[%v]<-->", cur.val)
		cur = cur.next
	}
	fmt.Println("<-tail")
}
