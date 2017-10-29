package main

import "fmt"

func main() {

	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"Dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
		"God morgon!",
	}

	for i , currentEntry := range greeting {
		fmt.Println(i, currentEntry)
	}

	for i := 0; i < len(greeting); i++ {
		fmt.Println(greeting[i])
	}
}
