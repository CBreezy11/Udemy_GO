package main

import (
	"fmt"
	"strconv"
)

func main() {

	c := inc(5)
	var count int

	for n := range c {
		count++
		fmt.Println(n)
	}
	fmt.Println("Final Count", count)
}

func inc(n int) <-chan string {
	out := make(chan string)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func(i int) {
			for k := 0; k < 20; k++ {
				out <- fmt.Sprint("Process:"+strconv.Itoa(i)+" Printing: ", k)
			}
			done <- true
		}(i)
	}
	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(out)
	}()
	return out
}
