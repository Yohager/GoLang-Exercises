package main 

import "fmt"


// 当连续的两个或者多个函数的已命名形参类型相同时，除最后一个类型外其他都可以忽略
// x int, y int可以缩写为 x, y int 
func add(x, y int) int {
	return x + y 
}

func main() {
	fmt.Println(add(100,124))
}