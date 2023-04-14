package main

import "fmt"

func main() {
	// 함수 리터럴을 사용하여 더하기 함수를 정의하고 호출합니다.
	sum := func(a, b int) int {
		return a + b
	}
	fmt.Println(sum(5, 3)) // Output: 8

	// 함수 리터럴을 인자로 전달하여 사용합니다.
	result := operate(5, 3, func(a, b int) int {
		return a * b
	})
	fmt.Println(result) // Output: 15
}

// 함수 리터럴을 인자로 받는 함수
func operate(a, b int, op func(int, int) int) int {
	return op(a, b)
}
