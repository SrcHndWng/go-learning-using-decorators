package main

import "fmt"

type Adder interface {
	Add(x, y int) int
}

type AdderFunc func(x, y int) int

func (a AdderFunc) Add(x, y int) int {
	return a(x, y)
}

func Do(adder Adder) int {
	return adder.Add(1, 2)
}

func main() {
	a := AdderFunc(
		func(x, y int) int {
			return x + y
		},
	)
	result := Do(a)
	fmt.Println(result)
}
