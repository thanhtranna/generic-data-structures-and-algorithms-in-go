package main

import (
	"fmt"
	"math"
	"time"
)

const size = 100_000

type Ordered interface {
	~float64 | ~int | ~string
}

func quickSort[T Ordered](data []T, low, high int) {
	if low < high {
		var pivot = partition(data, low, high)
		quickSort(data, low, pivot)
		quickSort(data, pivot+1, high)
	}
}

func partition[T Ordered](data []T, low, high int) int {
	// Pick a lowest bound element as a pivot value
	var pivot = data[low]
	var i = low
	var j = high

	for i < j {
		for data[i] <= pivot && i < high {
			i++
		}

		for data[j] > pivot && j > low {
			j--
		}

		if i < j {
			data[i], data[j] = data[j], data[i]
		}
	}

	data[low] = data[j]
	data[j] = pivot

	return j
}

func bubbleSort[T Ordered](data []T) {
	n := len(data)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func main() {
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = math.Sin(float64(i * i))
	}

	start := time.Now()
	quickSort[float64](data, 0, len(data)-1)
	elapsed := time.Since(start)
	fmt.Println("Elapsed sort time for sine wave using quicksort: ", elapsed)

	data = make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = math.Sin(float64(i * i))
	}

	start = time.Now()
	bubbleSort[float64](data)
	elapsed = time.Since(start)
	fmt.Println("Elapsed sort time for sine wave using bubblesort: ", elapsed)
}
