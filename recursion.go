package main

import (
	"fmt"
)

//1 1 2 3 5 8 13 21 34 55...

func fibSlice(time int) { //斐波那契数列非递归实现方式(Slice)

	fib := make([]int, 2)
	fib[0] = 1
	fib[1] = 1
	for i := 2; i < time; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	fmt.Println("斐波那契数列非递归实现方式(Slice):", fib)

}

func fibNormal(time int) { //斐波那契数列非递归实现方式(常规方法)

	a, b := 1, 1
	fmt.Print("斐波那契数列非递归实现方式(常规方法):")
	for i := 0; i < time; i++ {
		c := b
		b = a + b
		fmt.Print(a, " ")
		a = c
	}
}

func fibRe(time int) int { //斐波那契数列递归实现方式

	if time == 0 {
		return 0
	} else if time == 1 {
		return 1
	} else if time >= 2 {
		return fibRe(time-1) + fibRe(time-2)
	} else {
		return -1
	}

}

func fibClo() func() int { //斐波那契数列闭包实现方式

	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}

}

func main() {

	fibSlice(10)
	fibNormal(10)

	fmt.Println()
	fmt.Println("斐波那契数列递归实现方式:", fibRe(10))
	f := fibClo()

	fmt.Print("斐波那契数列闭包实现方式:")
	for i := 0; i < 10; i++ {
		fmt.Print(f(), " ")
	}

}
