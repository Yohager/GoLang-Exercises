package main 

import "fmt"

func add(x int, y int) {
	return x + y 
}

// 类型在变量名之后

func main() {
	fmt.Println(add(100,124))
}