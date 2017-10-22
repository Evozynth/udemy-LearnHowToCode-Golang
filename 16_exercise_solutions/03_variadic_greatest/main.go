package main

import "fmt"

func main() {
	fmt.Println(max(-88, -52, 3, -102, -99))
}

func max(numbers ...int) int {
	var largest int
	for i, v := range numbers {
		if v > largest || i == 0 {
			largest = v
		}
	}
	return largest
}
