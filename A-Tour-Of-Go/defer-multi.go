package main

import "fmt"

// 推迟的函数调用会被压入一个栈中
// 当外层函数返回时，被推迟的函数会按照后进先出的顺序调用

func main() {
	fmt.Println("Counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}
