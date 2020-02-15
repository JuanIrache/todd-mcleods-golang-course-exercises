package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	c := 0
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			t := c
			t++
			runtime.Gosched()
			c = t
			fmt.Printf("%d ", c)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("\nc reached %d\n", c)
}
