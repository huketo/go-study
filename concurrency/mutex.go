package main

import (
	"fmt"
	"sync"
	"time"
)

func philosopher(leftFork, rightFork *sync.Mutex) {
	for {
		leftFork.Lock()
		rightFork.Lock()
		fmt.Println("eating")
		leftFork.Unlock()
		rightFork.Unlock()
		fmt.Println("thinking")
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	forks := make([]*sync.Mutex, 5)
	for i := 0; i < 5; i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < 5; i++ {
		go philosopher(forks[i], forks[(i+1)%5])
	}

	var input string
	fmt.Scanln(&input)
}
