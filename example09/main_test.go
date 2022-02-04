package main

import (
	"sync"
	"testing"
)

func BenchmarkDivisors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divisors(30)
	}
}

func BenchmarkDivisors2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divisors2(30)
	}
}

func Divisors(n int) int {
	// Put your code here
	var counter int = 1
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			counter++
		}
	}
	return counter
}

// ------------------------------------------------------------------------------------------------------

func check(i int, n int, c chan int, wg *sync.WaitGroup) {
	if n%i == 0 {
		c <- 1
	}
	wg.Done()
}

func monitorWorker(wg *sync.WaitGroup, c chan int) {
	wg.Wait()
	close(c)
}

func Divisors2(n int) int {
	wg := &sync.WaitGroup{}
	c := make(chan int)
	d := 1

	for i := 1; i <= n/2; i++ {
		wg.Add(1)
		go check(i, n, c, wg)
	}

	go monitorWorker(wg, c)

	for o := range c {
		d = d + o
	}

	return d
}
