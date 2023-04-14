package main

import "fmt"

// 동물 인터페이스
type Animal interface {
	Speak() string
}

// 강아지 타입
type Dog struct {
	name string
}

// 강아지 타입이 Animal 인터페이스를 구현
func (d Dog) Speak() string {
	return "멍멍!"
}

// 고양이 타입
type Cat struct {
	name string
}

// 고양이 타입이 Animal 인터페이스를 구현
func (c Cat) Speak() string {
	return "야옹!"
}

func main() {
	// 인터페이스를 사용하여 다양한 동물들의 동작을 호출
	animals := []Animal{Dog{name: "뽀삐"}, Cat{name: "나비"}}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
