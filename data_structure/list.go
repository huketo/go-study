package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("List A")

	listA := list.New()

	listA.PushBack("A")
	listA.PushBack(100)
	listA.PushBack(true)
	// "A", 100, true
	for e := listA.Front(); e != nil; e = e.Next() {
		fmt.Printf("[%T] %v\n", e.Value, e.Value)
	}

	fmt.Println("-------------")
	fmt.Println("List B")

	listB := list.New()

	listB.PushBack("B")
	listB.PushFront("C")
	// "C", "B"
	for e := listB.Front(); e != nil; e = e.Next() {
		fmt.Printf("[%T] %v\n", e.Value, e.Value)
	}

	fmt.Println("-------------")
	fmt.Println("List A + List B")

	listA.PushBackList(listB)
	for e := listA.Front(); e != nil; e = e.Next() {
		fmt.Printf("[%T] %v\n", e.Value, e.Value)
	}
}
