package main

import "fmt"

func main() {
	x := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice Cream", "Sunsets"},
	}
	for k, v := range x {
		fmt.Println("This is the record for", k)
		for w, k2 := range v {
			fmt.Println("\t", w, k2)
		}
	}

}
