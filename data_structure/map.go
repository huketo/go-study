package main

import "fmt"

func main() {
	// 빈 맵 생성
	m := make(map[string]int)

	// key-value 쌍 추가
	m["apple"] = 1
	m["banana"] = 2
	m["orange"] = 3

	// 특정 key에 대한 value 조회
	fmt.Println("apple:", m["apple"])

	// 존재하지 않는 key에 대한 조회는 zero value 반환
	fmt.Println("grape:", m["grape"])

	// map 길이 출력
	fmt.Println("length:", len(m))

	// key-value 삭제
	delete(m, "orange")
	fmt.Println("after deletion:", m)
}
