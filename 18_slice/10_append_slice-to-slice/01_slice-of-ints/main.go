package main

import "fmt"

func main() {

	mySlice := []int{1, 2, 3, 4, 5}
	myOtherSlice := []int{6, 7, 8, 9}

	// Notice the syntax args...
	mySlice = append(mySlice, myOtherSlice...)

	fmt.Println(mySlice)

}
