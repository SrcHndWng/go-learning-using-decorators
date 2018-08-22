package main

import (
	"fmt"
	"log"
)

func decorator(f func(x, y int) int) func(x, y int) int {
	return func(x, y int) int {
		var result int
		// add log print
		defer func() {
			log.Printf("x=%v, y=%v, result=%v", x, y, result)
		}()
		result = f(x, y)
		return result
	}
}

func adderFunc(x, y int) int {
	// function imprement
	return x + y
}

func main() {
	result := decorator(adderFunc)(1, 2)
	fmt.Printf("result = %v\n", result)
}
