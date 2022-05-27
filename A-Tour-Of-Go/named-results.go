package main 

import "fmt"

// golang的返回值可以被命名 他们会被视作在函数顶部的变量
// 返回值的名称应当具有一定的意义，可以作为文档使用
// 没有参数的return语句返回已命名的返回值也就是直接返回
// 直接返回语句应当仅用在下面的这样的短函数中，长函数中它们会影响函数的可读性

func split(sum int) (x, y int) {
	x = sum * 4 / 9 
	y = sum - x 
	return 
}

func main() {
	fmt.Println(split(17))
}