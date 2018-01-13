package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 33
	fmt.Printf("%T %v \n", myAge, myAge)
}
