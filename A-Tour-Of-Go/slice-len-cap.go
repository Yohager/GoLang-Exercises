package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	// 截取切片使其长度为0
	s = s[:0]
	printSlice(s)
	// 拓展其长度
	s = s[:4]
	printSlice(s)
	// 舍弃前两个值
	s = s[2:]
	printSlice(s)
	// 额外扩展会报错
	// s = s[:10]
	// printSlice(s)
	// fmt.Println(s)

}
