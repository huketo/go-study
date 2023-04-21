package main

import (
	"fmt"
	"time"
)

func printNumbers(from, to int) {
	for i := from; i <= to; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go printNumbers(0, 4)       // 새로운 고루틴 생성
	go printNumbers(5, 9)       // 새로운 고루틴 생성
	time.Sleep(2 * time.Second) // main 고루틴 대기
}
