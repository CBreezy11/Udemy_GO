package main

import "fmt"

func main() {
	for bd := 1983; bd < 2020; bd++ {
		if bd%2 == 0 {
			fmt.Printf("%v Is an even year birthday\n", bd)
		} else {
			fmt.Printf("%v Is an odd year birthday", bd)
		}
	}
}
