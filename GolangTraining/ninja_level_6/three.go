package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("MainFucntion")
}

func foo() {
	defer func() {
		fmt.Println("First in foo")
	}()
	fmt.Println("last but first")
}
