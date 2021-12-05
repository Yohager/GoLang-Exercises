package main

import "os"

func main() {
	panic("a problem.")

	_,err := os.Create("/files/error.txt")
	if err != nil {
		panic(err)
	}
}
