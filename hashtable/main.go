package main

import (
	"fmt"
)

const ArraySize = 50

type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key string
	next *bucketNode
}

func main() {
	testHashTable := &HashTable{}
	fmt.Println(testHashTable)
}