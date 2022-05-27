package main

import "fmt"

// := 赋值可在类型明确的地方替代var声明
// 函数外的每个语句都必须以关键字开始(var, func等等)
// := 结构不能在函数外面使用

func main() {
	var i, j int = 1, 2
	i1, j1 := 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, i1, j1, k, c, python, java)
}
