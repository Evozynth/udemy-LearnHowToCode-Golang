package main

import "fmt"

func wrapper() func() int {
	x := 0
	// anonymous function (no name)
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}
