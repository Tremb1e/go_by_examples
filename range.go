package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	arr := []int{1, 2, 3, 4, 5}
	for _, i := range arr {
		fmt.Println(i)
	}

	//slice1 := make([]int, 5)
	slice1 := []int{2, 3, 4, 5, 6}
	for _, i := range slice1 {
		fmt.Println(i)
	}

	fmt.Println(slice1[2:5])

	map1 := make(map[string]int)
	map1["k1"] = 1
	map1["k2"] = 2
	map1["k3"] = 3
	//解决map遍历无序的问题
	var keys []string
	for k := range map1 {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, i := range keys {
		fmt.Println(i + "=" + strconv.Itoa(map1[i]))
	}

}
