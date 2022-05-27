package main

import "fmt"

// 注意和c以及java不一样的点
// go的for语句后面的三个构成部分外没有小括号同时大括号{}是必须的

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
