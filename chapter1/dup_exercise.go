package main

import (
	"bufio"
	"fmt"
	"os"
)

func judge(f *os.File, counts map[string]int) int {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	for _, n := range counts {
		if n > 1 {
			return 1
		}
	}
	return 0
}


func main() {
	//counts := make(map[string]int)
	files := os.Args[1:]
	for index,arg := range files {
		counts := make(map[string]int)
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dupExercise %v\n",err)
			continue
		}
		if judge(f,counts) == 1 {
			fmt.Printf("File %s has dup lines\n",files[index])
		}
		f.Close()
	}
}
