package main

import "fmt"

/*
golang的指针
类型 *T是指向T类型值的指针，其零值为nil
var p *int

& 操作符会生成一个指向其操作数的指针

* 操作符表示指针指向的底层值
这就是常说的间接引用或者说重定向

与C语言不同，golang没有指针运算
*/

func main() {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
