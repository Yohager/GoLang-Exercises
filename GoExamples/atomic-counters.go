package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var op uint64
	var cmp uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&op, 1)
				cmp += 1
			}

		}()
	}
	wg.Wait()

	fmt.Println("ops:", op)
	fmt.Println("cmp:", cmp)
	//不使用原子计数器会出现多个协程相互干扰使得运行的值出现改变
}
