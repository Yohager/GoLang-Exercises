package main

import "fmt"

// 切片的零值为nil 长度和容量都是0，且没有底层数组

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
