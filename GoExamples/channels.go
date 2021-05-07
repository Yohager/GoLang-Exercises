package main

import "fmt"

func main() {
	//使用make(chan val-type)创建一个新的通道，通道类型是传递值的类型
	messages := make(chan string)
	//使用channel <- 语法 发送一个新的值到通道中
	//这里发送的是"ping"
	go func() { messages <- "ping" }()
	//使用 <- channel 语法从通道中接受一个值，收到上面的ping消息并且打印出来
	//默认发送方和接收操作是阻塞的，直到发送方和接收方都就绪
	//这个特性允许我们不使用任何其他的同步操作就可以在程序结尾处等待消息ping
	msg := <-messages
	fmt.Println(msg)
}
