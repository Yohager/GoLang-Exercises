package main

import "fmt"

// 在声明一个变量而不指定其类型时（使用不带类型的:=或者var = ）
// 变量的类型由右值推导得出
// 当右值声明了类型时新变量的类型与其相同

func main() {
	v := 42.3
	j := v
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("j is of type %T\n", j)

}
