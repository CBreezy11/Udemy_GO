package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	inc := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mutex.Lock()
			v := inc
			v++
			inc = v
			fmt.Println(inc)
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("End Value", inc)

}
