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
	x := map[string]person{
		a.last: a,
		b.last: b,
	}
	for _, v := range x {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, j := range v.fav {
			fmt.Println(i, j)
		}
		fmt.Println("----------------------")
	}
}
