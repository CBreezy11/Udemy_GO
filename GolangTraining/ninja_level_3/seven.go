package main

import "fmt"

func main() {
	for bd := 1983; bd < 2020; bd++ {
		if bd%2 == 0 {
			fmt.Printf("%v Is an even year birthday\n", bd)
		} else if bd%5 == 0 {
			fmt.Printf("%v Is a year divisible by 5\n", bd)
		} else {
			fmt.Printf("%v Is an odd year birthday\n", bd)
		}
	}
}
