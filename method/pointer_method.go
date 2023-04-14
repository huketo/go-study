package main

import "fmt"

type Person struct {
	name string
	age  uint
}

// Value Receiver Methods
func (p Person) ChangeName_A(newName string) {
	p.name = newName
}

// Pointer Methods
func (p *Person) ChangeName_B(newName string) {
	p.name = newName
}

func main() {
	personA := &Person{name: "Jinsu", age: 23}
	fmt.Println("[0] personA:", personA)

	personA.ChangeName_A("Minsu")
	fmt.Println("=> ChangeName_A() : Value Receiver Methods")
	fmt.Println("[1] personA:", personA)

	personA.ChangeName_B("Minsu")
	fmt.Println("=> ChangeName_B() : Pointer Methods")
	fmt.Println("[2] personA:", personA)
}
