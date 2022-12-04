package main

import (
	"fmt"
	"time"

	"chapter05/nodestack"
	"chapter05/slicestack"
)

const size = 10_000_000

func main() {
	nodeStack := nodestack.Stack[int]{}
	sliceStack := slicestack.Stack[int]{}

	// Benchmark nodeStack
	start := time.Now()
	for i := 0; i < size; i++ {
		nodeStack.Push(i)
	}
	elapsed := time.Since(start)
	fmt.Println("\nTime for 10 million Push() operations on nodeStack: ", elapsed)
	start = time.Now()
	for i := 0; i < size; i++ {
		nodeStack.Pop()
	}
	elapsed = time.Since(start)
	fmt.Println("\nTime for 10 million Pop() operations on nodeStack: ", elapsed)
	// Benchmark sliceStack start = time.Now()
	for i := 0; i < size; i++ {
		sliceStack.Push(i)
	}

	elapsed = time.Since(start)
	fmt.Println("\nTime for 10 million Push() operations on sliceStack: ", elapsed)
	start = time.Now()
	for i := 0; i < size; i++ {
		sliceStack.Pop()
	}
	elapsed = time.Since(start)
	fmt.Println("\nTime for 10 million Pop() operations on sliceStack: ", elapsed)
}
