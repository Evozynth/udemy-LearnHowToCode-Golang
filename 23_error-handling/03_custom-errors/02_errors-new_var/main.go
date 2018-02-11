package main

import (
	"errors"
	"math"
	"log"
	"fmt"
)

var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func main() {
	fmt.Printf("%T\n", ErrNorgateMath)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorgateMath
	}
	// implementation
	return math.Sqrt(f), nil
}

// see use of errors.New in standard library:
// http://golang.org/src/bufio/bufio.go
// http://golang.org/src/pkg/io/io.go
