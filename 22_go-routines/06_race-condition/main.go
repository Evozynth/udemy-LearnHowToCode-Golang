package main

import (
	"sync"
	"fmt"
	"time"
	"math/rand"
)

var wg sync.WaitGroup
var counter int

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}

// go run -race main.go (to see if you have a race condition)
// vs
// go run main.go
