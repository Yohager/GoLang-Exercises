package main

import (
	"fmt"
)

//定义常量
const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g °F or %g °C\n",f,c)
}
