package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")
	//使用go f(s)在一个协程中调用这个函数，这个新的go协程会并发地执行这个函数
	go f("goroutine")
	//这两个协程在独立的协程中异步地运行，然后等待两个协程完成
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
