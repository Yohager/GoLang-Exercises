package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	//limiter通道每200ms接受一个值
	limiter := time.Tick(200 * time.Millisecond)
	//通过每次请求前阻塞limiter通道的一个接收
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
	//允许接受的爆发为3
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	//每200ms尝试添加一个新值到burstLimter中，上限为3
	go func() {
		for t:= range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()
	//下面模拟5个传入的请求，受益于burstLimiter的爆发前3个请求可以快速完成
	burstRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstRequests <- i
	}

	close(burstRequests)
	for req := range burstRequests {
		<- burstyLimiter
		fmt.Println("request",req,time.Now())
	}
}
