package main

import "fmt"

func main() {

	c := make(chan int)

	inc(c)
	recieve(c)
}

func recieve(c <-chan int) {
	for n := range c {
		fmt.Println(n)
	}

}

func inc(c chan<- int) {
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
}
