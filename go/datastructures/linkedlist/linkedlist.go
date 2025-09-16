package linkedlist

import (
	"errors"
	"fmt"
)

// Append, Prepend, InsertAfter(node, value), Delete(value), Find(value).

var IndexOutOfBounds = errors.New("index out of bounds")

type Node[T comparable] struct {
	next *Node[T]
	prev *Node[T]
	val  T
}

type LinkedList[T comparable] struct {
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

func (list *LinkedList[T]) Delete(value T) bool {
	if list.length == 0 {
		return false
	}

	// Find the node with the provided value
	cur := list.head
	for cur.val != value {
		cur = cur.next
	}

	// Value was not found
	if cur == nil {
		return false
	}

	// Last node in the list
	if cur.next == nil {
		previous := cur.prev
		cur.prev = nil
		previous.next = nil
		list.tail = previous
		// First node in the list
	} else if cur.prev == nil {
		next := cur.next
		cur.next = nil
		next.prev = nil
		list.head = next
	} else {
		previous := cur.prev
		next := cur.next
		cur.next = nil
		cur.prev = nil
		next.prev = previous
		previous.next = next
	}

	list.length--
	return true
}

func (list LinkedList[T]) Print() {
	fmt.Print("head->")
	cur := list.head
	for cur.next != nil {
		fmt.Printf("[%v]<-->", cur.val)
		cur = cur.next
	}
	fmt.Printf("[%v]<-tail\n", cur.val)
}
