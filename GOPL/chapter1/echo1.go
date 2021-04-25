package main

import (
	"fmt"
	"os"
	//"strconv"
)

func main() {
	//var s string
	//for i:=1; i<len(os.Args); i++{
	//	var s string = string(i) + os.Args[i]
	//	fmt.Print(s)
	//}
	for arg1,arg2 := range os.Args[1:] {
		//arg1 = strconv.Itoa(arg1)
		fmt.Println(arg1+1,arg2)
	}
}

