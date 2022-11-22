package main

import (
	"fmt"
	"time"
)

const LargestPrime = 100_000

var primes []int

// Send the sequence 3, 5, ... to channel 'ch'
func Generate(prime chan<- int) {
	for i := 3; ; i += 2 {
		prime <- i // Send 'i' to channel prime.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

func main() {
	start := time.Now()
	prime1 := make(chan int) // create a new channel.
	go Generate(prime1)      // Launch goroutine.

	for {
		prime := <-prime1 // Take prime1 out of channel
		if prime > LargestPrime {
			break
		}
		primes = append(primes, prime)
		prime2 := make(chan int)
		go Filter(prime1, prime2, prime)
		prime1 = prime2
	}

	elapsed := time.Since(start)
	fmt.Println("Computation time: ", elapsed)
	fmt.Println("Number of primes = ", len(primes))
}
