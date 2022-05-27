package main

import (
	"fmt"
	"math"
)

// the same as for
// 我们可以在if的条件表达式之前执行一个简单的语句
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 10),
	)
}
