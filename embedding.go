package main

import "fmt"

type personEmbedding struct {
	name string
	age  int
}

type container struct {
	personEmbedding
	color string
}

func (p personEmbedding) printperson() string {
	return fmt.Sprintf("Person的name是:%v", p.name)
}

func main() {
	co := container{
		personEmbedding: personEmbedding{
			name: "lisa",
			age:  18,
		},
		color: "yellow",
	}
	fmt.Println(co.name, co.age, co.color)
	fmt.Println(co)

	type printperson interface {
		printperson() string
	}

	var p printperson = co

	fmt.Println(p.printperson())
}
