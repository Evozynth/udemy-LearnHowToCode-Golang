package main

import (
	"net/http"
	"log"
	"bufio"
	"fmt"
)


func main() {
	const nrOfBuckets = 12

	// get the book adventures of sherlock holmes
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	// Set the split function fot the scanning operation.
	scanner.Split(bufio.ScanWords)

	// Create slice of slice of string to hold slices of words
	buckets := make([][]string, nrOfBuckets)

	// Loop over the words
	for scanner.Scan() {
		word := scanner.Text()
		n := HashBucket(word, nrOfBuckets)
		buckets[n] = append(buckets[n], word)
	}

	// Print len of each bucket
	for i := 0; i < nrOfBuckets; i++ {
		fmt.Println(i, "-", len(buckets[i]))
	}

	// Print the words in one of the buckets
	fmt.Println(buckets[6])
}

func HashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
