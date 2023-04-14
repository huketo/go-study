package main

import "fmt"

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}
	mySlice := myArray[1:4]

	fmt.Println("myArray: ", myArray)
	fmt.Println("mySlice: ", mySlice)
}
