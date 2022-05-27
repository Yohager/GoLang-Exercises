package main

import "fmt"

/*
常量的声明于变量类似只不过使用的是const关键词
常量可以是字符，字符串，布尔值和数值
常量不能使用:= 语法声明
*/

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
