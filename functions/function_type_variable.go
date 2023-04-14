package main

import "fmt"

// 함수 시그니처를 정의합니다.
type Calculator func(int, int) int

// 더하기 함수
func add(a, b int) int {
	return a + b
}

// 빼기 함수
func subtract(a, b int) int {
	return a - b
}

func main() {
	// 함수 타입 변수를 선언하고 함수를 할당합니다.
	var calc Calculator
	calc = add
	fmt.Println(calc(5, 3)) // Output: 8

	calc = subtract
	fmt.Println(calc(5, 3)) // Output: 2
}
