package main

import (
	"fmt"
	"math"
)

// 与for类似，无需小括号，但是一定需要大括号{}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-5))
}
