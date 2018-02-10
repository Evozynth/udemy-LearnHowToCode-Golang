package main

import (
	"fmt"
	"time"
)

func main() {
	c := factorial(4)
	fmt.Println("During the time it takes to calculate " +
		"the factorial we can do other stuff here...")

	fmt.Println("Like printing random stuff, or..")

	for i := 0; i < 6; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println("Do #", i)
	}

	fmt.Println("... as I just did, print a bunch of numbers")
	fmt.Println("Total:", <-c)
	fmt.Println("Stuff after the receiving from a channel is put on hold. " +
		"Maybe the result of the calculation is needed for something else below")
}

func factorial(n int) <-chan int {
	out := make(chan int)
	go func() {
		total := 1;
		for i := n; i > 0; i-- {
			fmt.Println("A lot of work to do in factorial function for n =", i)
			time.Sleep(time.Millisecond * 200) // Simulate something taking a lot of time
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}

/*
CHALLENGE #1
-- Use goroutines and channels to calculate factorial

CHALLENGE #2
-- Why might you want to use goroutines and channels to calculate factorial?
---- Formulate your own answer, then post that answer to this discussion area: https://goo.gl/flGsyX
---- Read a few of the other answers at the discussion area to see the reasons of others
 */
