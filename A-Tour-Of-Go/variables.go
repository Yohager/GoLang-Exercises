package main 

import "fmt"


// var语句用于声明一个变量列表，和函数参数的列表一样类型在最后
// var语句可以出现在包或者函数级别
var c, python, java bool 

func main() {
	var i int 
	fmt.Println(i, c, python, java)
}