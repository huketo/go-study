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
	defer func() {
		if r := recover(); r != nil {
			log.Println("recovered from panic:", r)
		}
	}()
	num, err := divide(300, 0)
	if err != nil {
		panic(fmt.Errorf("error occurred: %w", err))
	}
	// panic: error occurred: cannot divide through zero
	//
	// goroutine 1 [running]:
	// main.main()
	//         /home/huketo/go-study/error/error.go:39 +0x2ae
	// exit status 2
	fmt.Println("Number:", num)
}
