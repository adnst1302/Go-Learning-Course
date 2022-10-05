package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var x int64 = 0
	grup := sync.WaitGroup{}
	//var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			grup.Add(1)
			for j := 1; j <= 100; j++ {
				//x = x + 1
				atomic.AddInt64(&x, 1)
			}
			grup.Done()
		}()
	}
	grup.Wait()
	fmt.Println("Counter =", x)
}
