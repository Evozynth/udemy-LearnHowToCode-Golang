package main

import "fmt"

func main() {
	age := 44
	fmt.Println(&age) // 0xc420014128

	changeMe(&age)

	fmt.Println(&age) // 0xc420014128
	fmt.Println(age) // 24
}

func changeMe(z *int) {
	fmt.Println(z) // 0xc420014128
	fmt.Println(*z) // 44

	*z = 24
	fmt.Println(z) // 0xc420014128
	fmt.Println(*z) // 24
}
