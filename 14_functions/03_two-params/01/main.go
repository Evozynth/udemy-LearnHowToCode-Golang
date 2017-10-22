package main

import "fmt"

func main() {
	greet("John", "Doe")
}

func greet(fname string, lname string) {
	fmt.Println("Hello", fname, lname)
}
