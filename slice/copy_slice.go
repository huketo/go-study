package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice1_copy := slice1 // 슬라이스 복사

	fmt.Println("slice1:", slice1)
	fmt.Println("slice1_copy:", slice1_copy)

	slice1_copy[0] = 99 // slice1_copy의 첫 번째 요소 수정
	fmt.Println("=> slice1_copy 값 변경")

	fmt.Println("slice1:", slice1)           // 출력: [99 2 3 4 5]
	fmt.Println("slice1_copy:", slice1_copy) // 출력: [99 2 3 4 5]

	slice1_new_copy := make([]int, len(slice1))
	copy(slice1_new_copy, slice1) // 슬라이스의 값을 복사하여 새로운 메모리에 저장

	slice1_new_copy[0] = 1 // slice1_new_copy의 첫 번째 요소 수정
	fmt.Println("=> slice1_new_copy 값 변경")

	fmt.Println("slice1:", slice1)                   // 출력: [99 2 3 4 5]
	fmt.Println("slice1_copy:", slice1_copy)         // 출력: [99 2 3 4 5]
	fmt.Println("slice1_new_copy:", slice1_new_copy) // 출력: [99 2 3 4 5]
}
