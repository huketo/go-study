package main

import "fmt"

// 가변인수 함수 예시
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	// 가변인수 함수 호출
	result1 := sum(1, 2, 3, 4, 5) // 1 + 2 + 3 + 4 + 5 = 15
	fmt.Println("Result 1:", result1)

	nums := []int{6, 7, 8, 9, 10}
	result2 := sum(nums...) // 슬라이스를 풀어서 인수로 전달
	fmt.Println("Result 2:", result2)
}
