package main

import (
	"fmt"
	"strconv"
)

func printandLenSlice(slice []string) {

	fmt.Println(slice)
	fmt.Println(len(slice))
}

func main() {

	slice1 := make([]string, 5)

	for i := 0; i < 5; i++ {
		slice1[i] = "slice" + strconv.Itoa(i)
	}
	printandLenSlice(slice1)
	//fmt.Println(slice1)

	//fmt.Println(len(slice1))

	slice1 = append(slice1, "slice5", "slice6")
	printandLenSlice(slice1)
	//fmt.Println(slice1)

	//fmt.Println(len(slice1))

	slice2 := make([]string, len(slice1))
	copy(slice2, slice1)
	printandLenSlice(slice2)

	fmt.Println(slice2[2:5])

}
