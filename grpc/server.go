package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var total int64
	sum := 0
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		sum += i
		go func(i int) {
			defer wg.Done()
			atomic.AddInt64(&total, int64(i))
		}(i)
	}
	wg.Wait()

	fmt.Printf("total:%d sum %d", total, sum)
}