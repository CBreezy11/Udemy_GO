package main

import "fmt"

func foo() int {
	return 11
}

func bar() (int, string) {
	return 12, "string"
}

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}
