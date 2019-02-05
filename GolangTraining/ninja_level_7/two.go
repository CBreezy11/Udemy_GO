package main

import "fmt"

type person struct {
	name    string
	age     int
	address string
}

func changeMe(p *person) {
	p.address = "123 main st"
}

func main() {
	p1 := person{
		name:    "bob",
		age:     32,
		address: "123 bum st",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
