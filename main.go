package main

import "fmt"

type Adder interface {
	Add(x, y int) int
}

type adderImpl struct{}

func (adderImpl) Add(x, y int) int {
	return x + y
}

func main() {
	var a Adder = adderImpl{}
	result := a.Add(1, 2)
	fmt.Println(result)
}
