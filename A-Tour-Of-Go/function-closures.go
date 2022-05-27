package main

import "fmt"

/*
go函数可以是一个闭包
闭包是一个函数值
引用其函数体外的变量
该函数可以访问并赋予其引用的变量的值
该函数与这些变量绑定在一起
adder返回一个闭包
每个闭包都被绑定在其各自的sum变量上
*/
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
