package main

import "fmt"

type wheel int

var x wheel = 11

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

}
