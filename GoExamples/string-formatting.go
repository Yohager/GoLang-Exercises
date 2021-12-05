package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1,2}
	fmt.Printf("%v\n",p)
	//如果值是一个结构体 %+v的格式化输出内容将包括结构体的字段名
	fmt.Printf("%+v\n",p)
	//%#v 根据Go语言输出值产生该值的源码
	fmt.Printf("%#v\n",p)
	fmt.Printf("%T",p)
	fmt.Printf("%t\n",true)
	fmt.Printf("%d\n",123)

	fmt.Printf("%b\n",14)
	fmt.Printf("%c\n",33)
	// %x提供了十六进制的编码
	fmt.Printf("%x\n",456)
	//一些浮点数的格式
	fmt.Printf("%f\n",78.9)
	fmt.Printf("%e\n",123400000.0)
	fmt.Printf("%E\n",123400000.0)
	fmt.Printf("%s\n","\"string\"")
	fmt.Printf("%q\n","\"string\"")
	fmt.Printf("%x\n","hex this")
	//输出指针的值
	fmt.Printf("%p\n",&p)
	//一些对齐的格式
	fmt.Printf("|%6d|%6d|\n",12,345)

	fmt.Printf("|%-6.2f|%-6.2f|\n",1.2,3.45)
	fmt.Printf("|%6s|%6s|\n","foo","b")

	fmt.Printf("|%-6s|%-6s|\n","foo","b")

	s := fmt.Sprintf("a %s","string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr,"an %s \n","error")
}
