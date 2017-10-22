package main

import "fmt"

var x int // The idiomatic way, which sets int to its zero value that is 0
// var x = 0

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
Closure helps us limit the scope of variables used by multiple functions.
Without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope.
 */
