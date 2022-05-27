package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)
	// for i:= 0; i < dy; i++ {
	// 	cur := make([]=int{}, dx,dx)
	// 	for j,v := range cur {
	// 		res[i].append(res[i],j)
	// 	}
	// }
	for x := range res {
		b := make([]uint8, dx)
		for y := range b {
			b[y] = uint8(x*y - 1)
		}
		res[x] = b
	}
	return res
}

func main() {
	pic.Show(Pic)
}
