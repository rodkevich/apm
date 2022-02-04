package main

import "runtime"

func onServer(i int) { println("S:", i) }
func onUser(i int)   { println("U:", i) }

func main() {
	fromServer, fromUser := make(chan int), make(chan int)
	var serverData, userInput int
	var ok bool

	go func() {
		fromServer <- 1
		fromUser <- 1
		close(fromServer)
		runtime.Gosched()
		fromUser <- 2
		close(fromUser)
	}()

	isRunning := true
	for isRunning {
		select {
		case serverData, ok = <-fromServer:
			if ok {
				onServer(serverData)
			} else {
				isRunning = false
			}

		case userInput, ok = <-fromUser:
			if ok {
				onUser(userInput)
			} else {
				isRunning = false
			}
		}
	}
	println("end")
}
