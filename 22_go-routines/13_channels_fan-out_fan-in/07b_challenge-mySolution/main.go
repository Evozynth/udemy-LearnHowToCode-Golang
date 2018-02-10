package main

import (
	"sync"
	"fmt"
)

func main() {
	ch := make([]<-chan int, 10)
	in := gen()

	// FAN OUT
	// Multiple function reading from the same channel until that channel is closed
	// Distribute work across multiple function (ten goroutines) that all read from in.
	for i := range ch {
		ch[i] = factorial(in)
	}

	// FAN IN
	// multiplex multiple channels onto a single channel
	// merge the channels in the slice of channels onto a single channel
	var y int
	for n := range mergeChannelSlices(ch) {
		y++
		fmt.Println(y, "\t", n)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10000; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1;
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func mergeChannelSlices(cs []<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch <-chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

/*
CHALLENGE #1
-- Change the code above to execute 1,000 factorial computations concurrently and in parallel.
-- Use the "fan out / fan in" pattern to accomplish this

CHALLENGE #2
-- While running the factorial computations, try to find how much of your resources are being used.
-- Post the percentage of your resource being used to this discussion: https://goo.gl/BxKnOL
*/
