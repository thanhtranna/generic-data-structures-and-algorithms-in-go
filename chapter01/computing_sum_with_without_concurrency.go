package main

import (
	"fmt"
	"sync"
	"time"
)

var output1 float64
var output2 float64
var output3 float64
var output4 float64
var wg sync.WaitGroup

func worker1() {
	defer wg.Done()

	var output []float64

	sum := 0.0
	for index := 0; index < 100_000_000; index++ {
		output = append(output, 89.6)
		sum += 89.6
	}
	output1 = sum
}

func worker2() {
	defer wg.Done()

	var output []float64

	sum := 0.0
	for index := 0; index < 100_000_000; index++ {
		output = append(output, 64.8)
		sum += 64.8
	}
	output2 = sum
}

func worker3() {
	defer wg.Done()

	var output []float64

	sum := 0.0
	for index := 0; index < 100_000_000; index++ {
		output = append(output, 956.8)
		sum += 956.8
	}
	output3 = sum
}

func worker4() {
	defer wg.Done()
	var output []float64
	sum := 0.0
	for index := 0; index < 100_000_000; index++ {
		output = append(output, 1235.8)
		sum += 1235.8
	}
	output4 = sum
}

func main() {
	wg.Add(8)

	// Compute time with no concurrent processing
	start := time.Now()

	worker1()
	worker2()
	worker3()
	worker4()
	elapsed := time.Since(start)
	fmt.Println("\nTime for 4 workers in series: ", elapsed)
	fmt.Printf("Output1: %f \nOutput2: %f \nOutput3: %f \nOutput4: %f\n",
		output1, output2, output3, output4)

	// Compute time with concurrent processing
	start = time.Now()
	go worker1()
	go worker2()
	go worker3()
	go worker4()
	wg.Wait()
	elapsed = time.Since(start)
	fmt.Println("\nTime for 4 workers in parallel: ", elapsed)
	fmt.Printf("Output1: %f \nOutput2: %f \nOutput3: %f \nOutput4: %f",
		output1, output2, output3, output4)
}
