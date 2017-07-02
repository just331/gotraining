package main

import "fmt"

func main() {

	n := hashBucket("GO", 12)
	fmt.Println(n)

}

func hashBucket(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets
	return bucket

}
