package main

import (
	"fmt"
	"time"
)

// Brute force
func main() {
	start := time.Now()

	var found int
	n := 0

	for {
		n++
		for x := 2; x < 21; x++ {
			if n%x != 0 {
				break
			} else if x == 20 {
				found = n
			}
		}

		if found > 0 {
			break
		}
	}

	fmt.Println(found)
	fmt.Println(time.Since(start))
}

/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/
