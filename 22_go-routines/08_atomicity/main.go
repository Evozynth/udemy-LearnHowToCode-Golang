package main

import (
	"sync"
	"fmt"
	"time"
	"math/rand"
	"sync/atomic"
	"runtime"
)

var wg sync.WaitGroup
var counter int64

func init() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GOMAXPROCS(1)
}

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		// race:
		//x := counter
		//x++
		time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
		//counter = x
		// no race:
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}

// go run -race main.go (to see if you have a race condition)
// vs
// go run main.go
