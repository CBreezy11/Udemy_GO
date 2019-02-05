package main

import "fmt"

func main() {
	x := make(map[string][]string)
	x[`bond_james`] = []string{"Shaken, not stirred", "Martinis", "Women"}
	x[`moneypenny_miss`] = []string{"James Bond", "Literature", "Computer Science"}
	x[`no_dr`] = []string{"Being evil", "Ice Cream", "Sunsets"}
	for k, v := range x {
		fmt.Printf("This is the record for %v\n", k)
		for i, k2 := range v {
			fmt.Printf("\t %v %s\n", i, k2)
		}
	}
}
