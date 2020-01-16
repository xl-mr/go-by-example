package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)

	go func() {
		time.Sleep(time.Second * 10)
		message <- "ping"
	}()

	fmt.Println("test1")

	msg := <- message

	fmt.Println("test2")

	fmt.Println(msg)
}
