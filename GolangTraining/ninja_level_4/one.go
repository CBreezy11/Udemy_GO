package main

import "fmt"

func main() {
	x := [5]int{11, 12, 13, 14, 15}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
