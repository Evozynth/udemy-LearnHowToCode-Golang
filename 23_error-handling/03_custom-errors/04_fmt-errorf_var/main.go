package main

import (
	"math"
	"log"
	"fmt"
)


func main() {
	_, err := sqrt(-10.54)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		ErrNorgateMath := fmt.Errorf("norgate math again: square root of negative number: %v", f)
		return 0, ErrNorgateMath
	}
	// implementation
	return math.Sqrt(f), nil
}
