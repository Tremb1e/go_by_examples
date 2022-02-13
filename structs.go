package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 18
	return &p

}

func main() {

	fmt.Println(person{"karl", 18})
	fmt.Println(newPerson("kate"))
	fmt.Println(*newPerson("kate"))

	alpha := person{
		name: "alpha",
		age:  20,
	}
	fmt.Println(alpha)
	alpha.age = 21
	fmt.Println(alpha)

	bravo := &alpha
	bravo.age = 22
	fmt.Println(alpha)
	fmt.Println(bravo)

}
