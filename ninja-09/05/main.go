package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var c int64
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&c, 1)
			//It would make sense to remove Gosched, otherwise we are getting the value later, but writing is fine
			runtime.Gosched()
			fmt.Printf("%d ", atomic.LoadInt64(&c))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("\nc reached %d\n", atomic.LoadInt64(&c))
}
