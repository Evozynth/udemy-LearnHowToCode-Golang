package main

import "fmt"

func main() {
	name := "Libra"
	fmt.Println(name) // Libra
	fmt.Println(&name) // 0xc42000e1d0

	changeMe(&name)

	fmt.Println(&name) // 0xc42000e1d0
	fmt.Println(name) // Rocky
}

func changeMe(z *string) {
	fmt.Println(z) // 0xc42000e1d0
	fmt.Println(*z) // Libra

	*z = "Rocky"

	fmt.Println(z) // 0xc42000e1d0
	fmt.Println(*z) // Rocky
}
