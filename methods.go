package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

func (r rect) area() int {

	return r.width * r.height

}

func (r rect) perim() int {

	return (r.width + r.height) * 2

}

func addadd(i int) int {
	return i * i
}

type addint struct { //定义结构体
	a int
}

func (m addint) addaddd(n int) int { //结构体方法引用
	return m.a + n
}

func main() {

	r := rect{
		width:  3,
		height: 4,
	}

	fmt.Println(r.area())
	fmt.Println(r.perim())

	fmt.Println(addadd(5))

	m := addint{a: 1}         //实例化结构体
	fmt.Println(m.addaddd(2)) //结构体方法调用
	//结构体调用是.调用，简单方法调用是()内输入
}
