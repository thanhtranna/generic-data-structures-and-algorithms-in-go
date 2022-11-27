package main

import "fmt"

// This type implicitly implements Counter ADT
type counter struct {
	count int
}

// Interface serves to expose public features of counter
// The attribute count is private
type Counter interface {
	increment()
	decrement()
	reset()
	getCount() int
}

func (c *counter) increment() {
	c.count += 1
}

func (c *counter) decrement() {
	if c.count > 0 {
		c.count -= 1
	}
}

func (c *counter) reset() {
	c.count = 0
}

func (c *counter) getCount() int {
	return c.count
}

func main() {
	myCounter := Counter(&counter{})
	// The only operations that can be performed on myCounter
	// are specified in the Counter interface
	myCounter.increment()
	myCounter.increment()
	myCounter.reset()
	myCounter.increment()
	myCounter.increment()
	myCounter.increment()
	myCounter.increment()
	myCounter.decrement()
	countValue := myCounter.getCount()
	fmt.Println(countValue)
}
