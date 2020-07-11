package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n",false,"omit trailing newline")
var sep = flag.String("s"," ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(),*sep))
	if !*n {
		fmt.Println()
	}
}

/*
flag.Bool会创建一个新的对应布尔型标志参数的变量。它有三个属性，第一个是命令行标志参数的名字，第二个时标志参数的默认值，最后是该标志参数对应的描述信息
flag.String函数将用于创建一个对应字符串类型的标志参数变量，参数名，默认值，描述信息
程序中：sep和n变量分别指向对应命令行标志参数变量的指针，因此需要使用*sep和*n来间接引用他们
程序运行之前需要在使用标志参数对应的变量之前先调用flag.Parse函数用于更新参数
*/
