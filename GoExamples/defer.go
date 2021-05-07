package main

import "fmt"

/*
每次调用defer语句都会将函数压栈，函数参数会被拷贝下来，当外层函数退出
defer按照定义的逆序执行，如果defer执行的函数为nil则会在最终调用函数的产生panic
*/
func main() {
	var whatever [3]struct{}

	for i := range whatever {
		defer func() {
			fmt.Println(i)
		}()
		fmt.Println(i)
	}
}
//defer还有很多的用处，以后遇到了再总结
