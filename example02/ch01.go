package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	c := generator(1, 2, 3, 4, 5, 6, 7)
	out := squareator(c)
	for range c {
		fmt.Println(<-out)
	}
}

func generator(ints ...int) <-chan int {
	out := make(chan int)

	go func() {
		for i := range ints {
			out <- i
		}
		close(out)
	}()

	return out
}

func squareator(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for i := range in {
			out <- i * i
		}
	}()

	return out
}
