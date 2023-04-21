package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// 길이 3인 ring 생성
	r := ring.New(3)
	for i := 1; i <= 3; i++ {
		r.Value = i
		r = r.Next()
	}
	// 링의 요소를 연결
	r.Next().Link(r.Prev())

	r.Do(func(a interface{}) { fmt.Println(a) })
	// 1, 2, 3
	fmt.Println("현재 노드:", r.Value) // 1
	fmt.Println("---------------")

	// 1만큼 시계방향 회전, - 음수는 반시계
	r = r.Move(1)
	r.Do(func(a interface{}) { fmt.Println(a) })
	// 2, 3, 1
	fmt.Println("현재 노드:", r.Value) // 2
}
