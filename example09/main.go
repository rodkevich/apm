package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go goOne(ch1)
	go goTwo(ch2)

	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		default:
		}
	}
}

func goOne(ch chan string) {
	ch <- "From goOne goroutine"
	close(ch)

}

func goTwo(ch chan string) {
	time.Sleep(time.Second * 1)
	ch <- "From goTwo goroutine"
	close(ch)
}
