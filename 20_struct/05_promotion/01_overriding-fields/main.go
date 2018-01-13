package main

import "fmt"

type Person struct {
	First string
	Last string
	Age int
}

type DoubeZero struct {
	Person
	First string
	LicenseToKill bool
}

func main() {
	p1 := DoubeZero{
		Person: Person{
			First: "James",
			Last: "Bond",
			Age: 20,
		},
		First: "Double Zero Seven",
		LicenseToKill: true,
	}

	p2 := DoubeZero{
		Person: Person{
			First: "Miss",
			Last: "MoenyPenny",
			Age: 20,
		},
		First: "If looks could kill",
		LicenseToKill: false,
	}

	fmt.Println(p1.First, p1.Person.First)
	fmt.Println(p2.First, p2.Person.First)

}
