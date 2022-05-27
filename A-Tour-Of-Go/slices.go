package main

import "fmt"

/*
切片与python中的切片结果类似
*/

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)
}
