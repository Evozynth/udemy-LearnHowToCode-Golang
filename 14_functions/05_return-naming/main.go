package main

import "fmt"

func main() {
	fmt.Println(greet("John ", "Doe"))
}

func greet(fname string, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}
