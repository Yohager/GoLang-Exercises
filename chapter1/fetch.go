package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)


func main() {
	for _, url := range os.Args[1:] {
		//使用http.Get函数创建http请求
		//如果请求过程没有出错，resp会得到一个访问的结果
		//resp的Body字段会返回一个可读的服务器响应流
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr,"fetch:%v\n",err)
			os.Exit(1)
		}
		//使用ioutil.ReadAll函数得到resp中的全部值将其保存在变量b中
		b, err := ioutil.ReadAll(resp.Body)
		//Close用于防止资源泄露
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n",url,err)
			os.Exit(1)
		}
		fmt.Printf("%s",b)
	}
}