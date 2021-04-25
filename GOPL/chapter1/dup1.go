package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	//note: ignore potential errors from input.Err()
	for line, n := range counts {
		fmt.Printf("%d\t%s\n",n,line)
	}
}