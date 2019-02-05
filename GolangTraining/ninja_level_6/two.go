package main

import "fmt"

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar(x2 []int) int {
	total := 0
	for _, v := range x2 {
		total += v
	}
	return total
}

func main() {
	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ni := foo(i...)
	nii := bar(ii)
	fmt.Println(ni, nii)
}
