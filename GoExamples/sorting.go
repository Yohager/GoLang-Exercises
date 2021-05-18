package main

import (
	"fmt"
	"sort"
)

func main() {
	//对字符串进行排序
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)
	//对int类型的数字进行排序
	ints := []int{7, 2, 4, 5}
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)
	//检测一个切片是否有序
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

}
