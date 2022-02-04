package main

import (
	"log"
)

func main() {

	one := ppOne("one")
	two := ppTwo(one)
	exit := PpThree(two)
	for i := range exit {
		switch i {
		case struct{}{}:
			log.Println("finish it:", i)
		default:
			log.Println("idk")
		}
	}
	// for range exit {
	// 	fmt.Println(exit)
	// }

}

func ppOne(some string) <-chan string {
	ch := make(chan string)

	go func() {
		ch <- some
	}()
	return ch
}

func ppTwo(one <-chan string) <-chan uint {
	out := make(chan uint)

	go func() {
		for i := range one {
			switch i {
			case "one":
				log.Println("sending:", i)
				out <- 1
			}
		}
	}()

	return out
}

func PpThree(in <-chan uint) <-chan struct{} {
	exit := make(chan struct{})

	go func() {
		for i := range in {
			if i == 1 {
				log.Println("got:", i)

				exit <- struct{}{}
				close(exit)
			}
		}

	}()

	return exit
}
