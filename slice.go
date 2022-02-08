package main

import (
	"fmt"
	"strconv"
)

func printandLen(slice []string) {

	fmt.Println(slice)
	fmt.Println(len(slice))
}

func main() {

	slice1 := make([]string, 5)

	for i := 0; i < 5; i++ {
		slice1[i] = "slice" + strconv.Itoa(i)
	}
	printandLen(slice1)
	//fmt.Println(slice1)

	//fmt.Println(len(slice1))

	slice1 = append(slice1, "slice5", "slice6")
	printandLen(slice1)
	//fmt.Println(slice1)

	//fmt.Println(len(slice1))

	slice2 := make([]string, len(slice1))
	copy(slice2, slice1)
	printandLen(slice2)

	fmt.Println(slice2[2:5])

}
