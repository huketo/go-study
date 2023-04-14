package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats

	// 초기 메모리 할당 상태 확인
	runtime.ReadMemStats(&m)
	fmt.Println("초기 HeapAlloc:", m.HeapAlloc, "bytes")

	// 길이가 5인 int형 슬라이스 생성
	s := make([]int, 5)
	fmt.Println("슬라이스 초기 상태:", s)

	// 슬라이스에 값 할당
	s[0] = 1
	s[1] = 2
	s[2] = 3
	s[3] = 4
	s[4] = 5
	fmt.Println("슬라이스에 값 할당 후:", s)
	runtime.ReadMemStats(&m) // 메모리 할당 상태 다시 읽어옴
	fmt.Println("슬라이스 할당 후 HeapAlloc:", m.HeapAlloc, "bytes")

	// 슬라이스의 길이 변경
	s = s[:3] // 길이를 3으로 줄임
	fmt.Println("슬라이스 길이 변경 후:", s)

	// 가비지 컬렉션을 통한 메모리 회수
	// 슬라이스의 길이가 줄어들면, 가비지 컬렉션은 더 이상 사용되지 않는 메모리를 회수
	// 슬라이스의 길이가 3으로 줄어들었기 때문에, 더 이상 사용되지 않는 2개의 int 메모리를 회수
	runtime.GC()             // 가비지 컬렉션 실행
	runtime.ReadMemStats(&m) // 메모리 할당 상태 다시 읽어옴
	fmt.Println("가비지 컬렉션을 통한 메모리 회수 후 HeapAlloc:", m.HeapAlloc, "bytes")

	// 메모리 할당 상태 확인
	fmt.Println("총 메모리 크기 HeapSys:", m.HeapSys, "bytes") // 힙 시스템의 총 메모리 크기 출력
}
