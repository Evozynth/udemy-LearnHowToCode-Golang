package main

import (
	"net/http"
	"log"
	"bufio"
	"fmt"
)


func main() {
	const nrOfBuckets = 12

	// get the book moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set the split function fot the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Create slice to hold counts
	buckets := make([]int, nrOfBuckets)
	// Loop over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text(), nrOfBuckets)
		buckets[n]++
	}
	fmt.Println(buckets)
}

func HashBucket(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}
