package main

import "fmt"

func add(nums ...int) {
	fmt.Print(nums, " ")
	add := 0
	for _, i := range nums {
		add += i
	}
	fmt.Println(add)

}

func printt(str ...string) {

	fmt.Print(str)

}

func main() {
	add(1, 2, 3)
	printt("test print")
}
