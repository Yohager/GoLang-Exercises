package main

import (
	"fmt"
	"math"
)

/*
表达式T(v)将值v转换为类型T
可以使用 var i int = int(x) 也可以使用 i := int(x)
golang中不同类型的项之间赋值需要显式转换
*/

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z = uint(f)
	fmt.Println(x, y, z, f)
	fmt.Printf("%T", z)
}
