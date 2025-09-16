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

	fmt.Println("Inserting node after index 4")
	err := list.InsertAfter(4, 6)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	list.Print()

	fmt.Println("Deleting node at the beginning")
	list.Delete(1)
	list.Print()
	fmt.Println("Deleting node in the middle")
	list.Delete(3)
	list.Print()
	fmt.Println("Deleting node at the end")
	list.Delete(6)
	list.Print()

}
