package main

import "fmt"

func main() {

	msg := make(chan string, 2)

	msg <- "msg1"
	msg <- "msg2"

	fmt.Println(<-msg)
	fmt.Println(<-msg)
}
