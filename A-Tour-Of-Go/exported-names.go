package main

import (
	"fmt"
	"math"
)

// golang中 如果一个名字以大写字母开头那么他就是已导出的
// Pi就是一个导自于math的包
// pi就是未导出的

func main() {
	fmt.Println(math.Log(2))
	fmt.Println(math.Pi)
}
