package main

import "fmt"

func zeroval(ival int) {
	ival = 1
}

func zeroptr(iptr *int) {
	*iptr = 2
}

func main() {

	i := 0
	fmt.Println(i)

	zeroval(i)
	fmt.Println(i)

	zeroptr(&i)
	fmt.Println(i)
	fmt.Println(&i)

	var k *int //地址转换，需要注意k的类型是*int
	j := 3
	k = &j //将j的地址放入k中
	fmt.Println(k)
	fmt.Println(*k) //输出时使用*，取出k中内容所对应内存位置的值
}
