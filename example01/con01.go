package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func one(a string) {
	fmt.Println(a)
	wg.Done()
}

func two(b string) {
	fmt.Println(b)
	wg.Done()
}

func main() {
	wg.Add(2)
	go one("string for one")
	go two("string for two")
	time.Sleep(3 * time.Second)
	wg.Wait()
}
