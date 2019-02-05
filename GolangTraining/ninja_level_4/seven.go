package main

import "fmt"

func main() {
	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(xs1)
	fmt.Println(xs2)
	xxs := [][]string{xs1, xs2}
	fmt.Println(xxs)
	for i, xs := range xxs {
		fmt.Printf("Record: %v\n", i)
		for j, val := range xs {
			fmt.Printf("\t Index Position: %v \t Value: %v\n", j, val)
		}
	}

}
