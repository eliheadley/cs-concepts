package main

import (
	"fmt"

	ll "cs-concepts/datastructures/linkedlist"
)

func main() {
	fmt.Println("Welcome to the Go Linked List Example")
	list := ll.LinkedList[int]{}

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	list.Print()
	err := list.InsertAfter(5, 6)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	list.Print()
}
