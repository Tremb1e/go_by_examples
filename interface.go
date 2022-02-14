package main

import "fmt"

type phone interface { //接口中定义了一个通用的方法
	call()
}

type Nokia struct {
	name string
}

type Sansung struct {
	name string
}

func (nokia Nokia) call() {
	nokia.name = "Nokia"
	fmt.Println(nokia.name, " call")
}

func (sansung Sansung) call() {
	sansung.name = "Sansung"
	fmt.Println(sansung.name, " call")
}

func main() {
	var nokia phone

	nokia = new(Nokia)
	nokia.call()

	var sansung phone
	sansung = new(Sansung)
	sansung.call()
}
