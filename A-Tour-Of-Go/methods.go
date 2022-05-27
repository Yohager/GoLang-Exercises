package main

import (
	"fmt"
	"math"
)

/*
golang中没有类，我们可以为结构体类型定义方法
方法是一类带特殊的接收者参数的函数
方法接收者在它自己的参数列表内，位于func关键字和方法名之间
在此例中，Abs方法拥有一个名为v类型为Vertex的接收者
*/

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
