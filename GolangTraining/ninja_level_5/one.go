package main

import "fmt"

type person struct {
	first string
	last  string
	fav   []string
}

func main() {
	a := person{"James", "Bond", []string{"chocolate", "martini", "rum and coke"}}
	b := person{"Miss", "Moneypenny", []string{"strawberry", "Vanilla", "Capuccino"}}
	fmt.Println(a.first)
	fmt.Println(a.last)
	for i, j := range a.fav {
		fmt.Println(i, j)
	}
	fmt.Println(b.first)
	fmt.Println(b.last)
	for i, j := range b.fav {
		fmt.Println(i, j)
	}
}
