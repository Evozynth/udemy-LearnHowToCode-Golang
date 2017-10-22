package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	// var xs []int // Will create a nil slice of integers (no memory address)
	xs := []int{} // Needs to be initialized manually, however it's not a nil slice but an actual empty slice with a memory address
	for _, v := range numbers {
		if callback(v) {
			xs = append(xs, v)
		}
	}
	return xs
}

func main() {
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs) // [2 3 4]
}
