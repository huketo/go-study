package main

import (
	"errors"
	"fmt"
	"log"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, errors.New("cannot divide through zero")
	}
	return a / b, nil
}

func main() {
	num1, err := divide(100, 0)
	if err != nil {
		fmt.Println("error occurred:", err)
	}
	// error occurred: cannot divide through zero
	fmt.Println("Number:", num1)

	fmt.Println("------------------------")

	num2, err := divide(200, 0)
	if err != nil {
		log.Printf("error occurred: %v", err)
	}
	// 2023/04/20 22:15:52 error occurred: cannot divide through zero
	fmt.Println("Number:", num2)
}
