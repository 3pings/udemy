package main

import (
	"log"
	"net/http"
	"bufio"
	"fmt"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	scanner.Split(bufio.ScanWords)
	buckets := make([]int, 12)
	for scanner.Scan() {
		n := HashBucket(scanner.Text(), 12)
		buckets[n]++
	}
	fmt.Println(buckets)

}

func HashBucket (word string, buckets int) int {
	// letter := rune(word[0])
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}