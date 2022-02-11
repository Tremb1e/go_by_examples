package main

import "fmt"

type Website struct { //定义网站对象
	name   string
	number int
}

func addClo() func() int { //闭包返回计数值

	i := 0
	return func() int {
		i++
		return i
	}

}

func (website Website) Printwebsite() {

	fmt.Println("网站名称", website.name)
	fmt.Println("第几个网站", website.number)

}

func main() {

	i := addClo()
	web1 := Website{ //实例化web1
		name:   "www.google.com",
		number: i(),
	}

	web2 := Website{ //实例化web2
		name:   "www.baidu.com",
		number: i(),
	}
	web1.Printwebsite() //web1输出两边证明闭包计数功能不同于i++
	web1.Printwebsite()
	web2.Printwebsite()
	web2.Printwebsite()

}
