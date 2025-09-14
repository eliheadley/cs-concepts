package main

import (
	"fmt"

	"datastructures/linkedlist"
)

func main() {
	fmt.Println("Welcome to the Go Linked List Example")
	list := linkedlist.LinkedList[int]{}
	list.Print()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	list.Print()
}
