package main

import (
	"errors"
	"fmt"
)

func adderror(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return -1, errors.New("不能比0小")
	}
	return a + b, nil
}

func main() {
	i, error := adderror(-1, 2)
	fmt.Println(i, error)
}
