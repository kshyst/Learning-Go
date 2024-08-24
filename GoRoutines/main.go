package main

import (
	"fmt"
	"time"
)

func main() {
	var messages = make(chan string)

	go func(message string) {
		for i := 0; i < 3; i++ {
			println(message+" ", i)
			messages <- message
		}
	}("first goroutine")

	fmt.Println("doneee")

	go func(message string) {
		for i := 0; i < 3; i++ {
			println(message+" ", i)
		}
	}("second goroutine")

	msg := <-messages
	fmt.Println(msg)
	time.Sleep(time.Second)
	fmt.Println("done")

	fmt.Println("----------------")

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	fmt.Println("----------------")

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		}
	}

	fmt.Println("----------------")

	//timeout implementation

	chan1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		chan1 <- "result 1"
	}()

	select {
	case res := <-chan1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	chan2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		chan2 <- "result 2"
	}()

	select {
	case res := <-chan2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
