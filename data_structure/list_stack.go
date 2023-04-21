package main

import (
	"container/list"
	"fmt"
)

func main() {
	stack := list.New()

	// Push elements onto stack
	stack.PushBack(1)
	stack.PushBack(2)
	stack.PushBack(3)

	// Pop elements from stack
	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)
		fmt.Println(e.Value)
	}
}
