package main

import (
	"fmt"
	"math/rand"
	"time"

	"chapter10/set/floatset"
)

const (
	size = 1_000_000
)

var dataSet []float64

func main() {
	mySet := floatset.NewSet()

	dataSet = make([]float64, size)
	for i := 0; i < size; i++ {
		dataSet[i] = 100.0 * rand.Float64()
	}
	// Time construction of Set
	start := time.Now()
	for i := 0; i < size; i++ {
		mySet.Add(dataSet[i])
	}
	elapsed := time.Since(start)
	fmt.Printf("\nTime to build Set with %d numbers: %s", size, elapsed)

	// Time to test the presence of all numbers in dataSet
	start = time.Now()
	for i := 0; i < len(dataSet); i++ {
		if !mySet.IsPresent(dataSet[i]) {
			fmt.Printf("%f not present\n", dataSet[i])
		}
	}
	elapsed = time.Since(start)
	fmt.Printf("\nTime to test the presence of all numbers in Set: %s", elapsed)

	avlSet := floatset.AVLTree{nil, 0}
	// Time construction of avlSet
	start = time.Now()
	for i := 0; i < size; i++ {
		avlSet.Insert(dataSet[i])
	}
	elapsed = time.Since(start)
	fmt.Printf("\n\nTime to build avlSet with %d numbers: %s", size, elapsed)

	// Time to test the presence of all numbers in avlSet
	start = time.Now()
	for i := 0; i < len(dataSet); i++ {
		if !mySet.IsPresent(dataSet[i]) {
			fmt.Printf("%f not present\n", dataSet[0])
		}
	}
	elapsed = time.Since(start)
	fmt.Printf("\nTime to test the presence of all numbers in avlSet: %s", elapsed)

	// Use concurrent processing to construct concurrent avl trees
	start = time.Now()
	floatset.BuildConcurrentSet(dataSet)
	elapsed = time.Since(start)
	fmt.Printf("\n\nTime to build concurrent (%d) avlSet with %d numbers: %s", floatset.Concurrent, size, elapsed)

	// Test every number in dataSet against the concurrent set
	start = time.Now()
	for i := 0; i < len(dataSet); i++ {
		if !floatset.IsPresent(dataSet[i]) {
			fmt.Printf("%f not present\n", dataSet[i])
		}
	}
	elapsed = time.Since(start)
	fmt.Printf("\nTime to test the presence of all numbers in concurrent (%d) avlSet: %s", floatset.Concurrent, elapsed)
}

/*
On iMac Pro with 16G Ram and 2.9 GHz 6-Core Intel Core i5

Time to build Set with 1000000 numbers: 146.798542ms
Time to test the presence of all numbers in Set: 88.283394ms

Time to build avlSet with 1000000 numbers: 712.453284ms
Time to test the presence of all numbers in avlSet: 88.8373ms

Time to build concurrent (32) avlSet with 1000000 numbers: 205.415125ms
Time to test the presence of all numbers in concurrent (32) avlSet: 66.387052ms

*/
