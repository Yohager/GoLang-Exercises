package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main()  {
	//是否包含某一个字符串
	p("Contains: ",s.Contains("test","es"))
	//统计某个字符出现的次数
	p("Count: ",s.Count("test","t"))
	//是否是以指定的前缀开头
	p("HasPrefix: ",s.HasPrefix("test","te"))
	//是否以指定的后缀结尾
	p("HasSuffix: ",s.HasSuffix("test","st"))
	//找第一次出现的索引，没有则返回-1
	p("Index: ",s.Index("test","e"))
	//字符串连接
	p("Join: ",s.Join([]string{"a","b"},"-"))
	//创建重复的字符串
	p("Repeat : ",s.Repeat("a",5))
	//替换功能
	p("Replace: ",s.Replace("foo","o","0",-1))
	//替换多少个
	p("Replace: ",s.Replace("foo","o","0",1))
	//字符串分割
	p("Split: ",s.Split("a-b-c-d-e","-"))
	//转小写
	p("ToLower: ",s.ToLower("TEST"))
	//转大写
	p("ToUpper: ",s.ToUpper("test"))
	p()
	//字符串长度
	p("Len: ",len("hello"))
	//给出指定索引位置的utf-8编码
	p("Char: ","hello"[0])
}
