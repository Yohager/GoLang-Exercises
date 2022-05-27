package main

import (
	"fmt"
	"math/cmplx"
)

/*
基本类型
bool string
int: int8, int16, int32, int64
uint: uint8, uint16, uint32, uint64, uintptr
byte: uint8的别名
rune: int32的别名 表示一个Unicode码点
float: float32, float64
complex: complex64, complex128
int, uint, uintptr在32位的系统中通常32位宽，在64位的系统中为64位宽
正常使用int就完事了

*/

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("type: %T value: %v\n", ToBe, ToBe)
	fmt.Printf("type: %T value: %v\n", MaxInt, MaxInt)
	fmt.Printf("type: %T value: %v\n", z, z)
}
