package main

import "fmt"

func main() {
	student := make([]string, 35)
	students := make([][]string, 35)
	student[0] = "Libra"
	// student = append(student, "Libra")
	fmt.Println(student)
	fmt.Println(students)
}
