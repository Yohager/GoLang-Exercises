package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	last := x
	eps := 0.0000001
	for i := 1; math.Abs(last-z) > eps; i++ {
		last = z
		z -= (z*z - x) / (2 * z)
		// fmt.Println(z)
	}
	return z

}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
