package main

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

// Rectangle 타입에 대한 메서드 정의
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	rectangle := Rectangle{width: 10, height: 5}
	fmt.Println("넓이:", rectangle.Area()) // 메서드 호출
}
