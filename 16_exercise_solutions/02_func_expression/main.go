package main

import "fmt"

func main() {

	half := func(x int) (int, bool) {
		return x/2, x%2 == 0
	}

	v := 1
	x, b := half(v);
	fmt.Printf("Half of %v is %v and is even? %v\n", v, x, b)

	v = 2
	x, b = half(v);
	fmt.Printf("Half of %v is %v and is even? %v\n", v, x, b)
}
