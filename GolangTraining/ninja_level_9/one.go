package main

import "fmt"

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println("Hello")

}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		name: "James",
	}
	p1.speak()
}
