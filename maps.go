package main

import "fmt"

func printandLenMap(x map[string]int) {
	fmt.Println(x)
	fmt.Println(len(x))
}

func main() {

	map1 := make(map[string]int)

	map1["key1"] = 1
	map1["key2"] = 2
	map1["key3"] = 3
	map1["key4"] = 4

	printandLenMap(map1)

	delete(map1, "key2")
	printandLenMap(map1)

	_, prs := map1["key2"]
	fmt.Println(prs)

	map2 := map[string]int{"test1": 111, "test2": 222}
	printandLenMap(map2)
}
