package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	start := time.Now()

	c := gen()

	f1 := factorial(c)
	f2 := factorial(c)
	f3 := factorial(c)
	f4 := factorial(c)
	f5 := factorial(c)
	f6 := factorial(c)
	f7 := factorial(c)
	f8 := factorial(c)
	f9 := factorial(c)
	f := factorial(c)

	var y int

	for x := range merge(f, f1, f2, f3, f4, f5, f6, f7, f8, f9) {
		y++
		fmt.Println(y, "\t", x)
	}

	fmt.Println(time.Since(start))

}

func merge(ch ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ch))

	for _, c := range ch {
		go func(ch <-chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)

		go func() {
			wg.Wait()
			close(out)
		}()
	}
	return (out)
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for x := range in {
			out <- fact(x)
		}
		close(out)
	}()
	return (out)
}

func fact(x int) int {
	total := 1
	for n := x; n > 0; n-- {
		total *= n
	}
	return total
}

/*
CHALLENGE #1:
-- Change the code above to execute 100 factorial computations concurrently and in parallel.
-- Use the "pipeline" pattern to accomplish this

POST TO DISCUSSION:
-- What realizations did you have about working with concurrent code when building your solution?
-- eg, what insights did you have which helped you understand working with concurrency?
-- Post your answer, and read other answers, here: https://goo.gl/uJa99G
*/
