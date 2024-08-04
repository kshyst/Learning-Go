package main

import "fmt"

func main() {

	var c = make(chan string)

	go dataInjection(c)
	go dataInjection2(c)

	for i := 0; i < 2; i++ {
		select {
		case msg := <-c:
			fmt.Println(msg)
		case msg2 := <-c:
			fmt.Println(msg2)
		}

	}
}

func dataInjection(c chan string) {
	c <- "Hello"
}

func dataInjection2(c chan string) {
	c <- "World"
}
