package main

import (
	"math"
	"log"
	"fmt"
)

type NorgateMathError struct {
	lat, long string
	err       error
}

func (n *NorgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.54)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		return 0, &NorgateMathError{"50.683281 N", "99.469509 W", nme}
	}
	// implementation
	return math.Sqrt(f), nil
}

// see use of structs with error type in standard library:
// http://goinggo.net/2014/11/error-handling-in-go-part-ii.html
//
// http://golang.org/pkg/ner/#OpError
// http://golang.org/src/pkg/net/dial.go
// http://golang/org/src/pkg/net/net.go
//
// http://golang.org/src/pkg/encoding/json/decode.go
