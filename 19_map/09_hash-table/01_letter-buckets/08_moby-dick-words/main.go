package main

import (
	"net/http"
	"log"
	"bufio"
	"fmt"
)

func main() {
	// get the book moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	// NewScanner takes a reader and res.Body implements the reader interface (so it is a reader)
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close();
	// Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	// Loop over the word
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
