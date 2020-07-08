package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func main() {
	//首先进行一个空映射
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			//正常读取则err返回nil，这种情况下未成功读取返回error
			fmt.Fprintf(os.Stderr,"dup3:%v\n",err)
			continue
		}
		for _, line := range strings.Split(string(data),"\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}
