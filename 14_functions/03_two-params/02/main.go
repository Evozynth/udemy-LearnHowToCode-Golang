package main

import "fmt"

func main() {
	greet("John", "Doe")
}

// Since both fname and lname is strings it's possible to skip
// the type declaration of the first parameter
func greet(fname , lname string) {
	fmt.Println("Hello", fname, lname)
}
