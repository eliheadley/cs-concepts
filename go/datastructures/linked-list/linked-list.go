package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) add(newVal int) {
	temp := Node{nil, newVal}
	if l.head.next == nil {
		l.head = &temp
		return
	}

	cur := *l.head
	for cur.next != nil {
		cur = *cur.next
	}
	cur.next = &temp
}

func (l LinkedList) print() {
	fmt.Print("\nhead->")
	cur := *l.head
	for cur.next != nil {
		fmt.Printf("%v->", cur.val)
		cur = *cur.next
	}
	fmt.Println("end")
}

func main() {
	list := LinkedList{}

	list.add(1)
	list.add(2)
	list.add(3)
	list.add(4)
	list.add(5)

	list.print()
}
