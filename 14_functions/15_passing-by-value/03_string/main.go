package main

import "fmt"

func main() {
	name := "Libra"
	fmt.Println(name) // Libra

	changeMe(name)

	fmt.Println(name) // Libra
}

func changeMe(z string) {
	fmt.Println(z) // Libra
	z = "Rocky"
	fmt.Println(z) // Rocky
}
