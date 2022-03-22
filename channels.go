package main

import "fmt"

func main() {
	messages := make(chan string)
	go func() { messages <- "channels" }()
	msg := <-messages
	fmt.Println(msg)
}
