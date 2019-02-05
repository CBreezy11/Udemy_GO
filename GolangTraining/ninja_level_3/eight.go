package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("Will not print")
	case true:
		fmt.Println("Will print")

	}
}
