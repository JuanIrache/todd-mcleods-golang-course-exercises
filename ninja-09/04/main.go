package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	c := 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			t := c
			t++
			//It would make sense to remove Gosched
			runtime.Gosched()
			c = t
			fmt.Printf("%d ", c)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("\nc reached %d\n", c)
}
