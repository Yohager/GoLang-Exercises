package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

/*
结构体的字段可以通过结构体指针来访问

如果我们有一个指向结构体的指针p
那么可以通过(*p).X 来访问其字段
golang是允许我们使用隐式间接引用直接写p.X即可
*/

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
