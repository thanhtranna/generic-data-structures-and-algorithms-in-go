package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

const size = 100_000_000

var data []float64

func isSorted1(data []float64) bool {
	var data1 []float64

	data1 = make([]float64, len(data))
	copy(data1, data) // Copies data into data1

	// Compare data and data1
	for i := 0; i < size; i++ {
		if data[i] != data1[i] {
			return false
		}
	}

	return true
}

func isSorted2(data []float64) bool {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}

func isSorted3(data []float64) bool {
	ch := make(chan bool)
	numSegments := runtime.NumCPU()
	fmt.Println("NumCPU", numSegments)
	segmentSize := int(float64(len(data)) / float64(numSegments))
	// Launch numSegments goroutines
	for index := 0; index < numSegments; index++ {
		go isSegmentSorted(data, index*segmentSize, index*segmentSize+segmentSize, ch)
	}

	num := 0 // completed goroutines
	for {
		select {
		case value := <-ch: // Blocks until a goroutines puts a bool into the channel
			if !value {
				return false
			}
			num += 1
			if num == numSegments { // All goroutines have completed
				return true
			}
		}
	}

	return true
}

func isSegmentSorted(data []float64, a, b int, ch chan<- bool) {
	// Generates boolean value put into ch
	for i := a + 1; i < b; i++ {
		if data[i] < data[i-1] {
			ch <- false
		}
	}

	ch <- true
}

func main() {
	data = make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = 100.0 * rand.Float64()
	}

	fmt.Println("add done data")

	data2 := make([]float64, size)
	for i := 0; i < size; i++ {
		data2[i] = float64(2 * i)
	}

	fmt.Println("add done data2")

	start := time.Now()
	result := isSorted1(data)
	elapsed := time.Since(start)
	fmt.Println("Sorted: ", result)
	fmt.Println("elapsed using sorted1:", elapsed)

	start = time.Now()
	result = isSorted1(data2)
	elapsed = time.Since(start)
	fmt.Println("Sorted: ", result)
	fmt.Println("elapsed using sorted1:", elapsed)

	start = time.Now()
	result = isSorted2(data)
	elapsed = time.Since(start)
	fmt.Println("\nSorted: ", result)
	fmt.Println("elapsed using sorted2", elapsed)

	start = time.Now()
	result = isSorted2(data2)
	elapsed = time.Since(start)
	fmt.Println("Sorted: ", result)
	fmt.Println("elapsed using sorted2:", elapsed)

	start = time.Now()
	result = isSorted3(data)
	elapsed = time.Since(start)
	fmt.Println("\nSorted: ", result)
	fmt.Println("elapsed using sorted3", elapsed)

	start = time.Now()
	result = isSorted3(data2)
	elapsed = time.Since(start)
	fmt.Println("Sorted: ", result)
	fmt.Println("elapsed using sorted3:", elapsed)
}
