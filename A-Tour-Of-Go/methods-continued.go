package main

import (
	"fmt"
	"math"
)

/*
可以为非结构体类型声明方法
我们看到这里带Abs方法的数值类型是MyFloat
我们只能为在同一个包内定义的类型的接收者声明方法
不能为其他包内定义的类型(包括int之类的内建类型)的接收者声明方法
(也就是说接收者的类型定义和方法必须在同一个包内，不能为内建的声明方法)

*/

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
