package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i < 5; i++ {
		time.Sleep(time.Millisecond * 250)
		fmt.Println(i)
	}
}

func alphabets() {
	for i := 'a'; i < 'e'; i++ {
		time.Sleep(time.Millisecond * 400)
		fmt.Println(i)
	}
}

func main() {
	go numbers()
	go alphabets()
	time.Sleep(time.Second * 3)
	fmt.Println("main")
}
