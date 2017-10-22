package main

import "fmt"

func visit(numbers []int, callback func(int)) {
	for _, v := range numbers {
		callback(v)
	}
}

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
}

// callback: passing a func as an argument
