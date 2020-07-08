package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

//color.Color是一种复合声明，生成的是一个切片
var palette = []color.Color{color.White,color.Black}

const (
	whiteIndex = 0 //first color in palette
	blackIndex = 1 //next color in palette
)

func lissajous(out io.Writer) {
	//这里的const常量只能在函数内部使用
	const (
		cycles        = 5       //number of complete x oscillator revolutions
		res           = 0.001   //angular resolution
		size, nframes = 100, 64 //image canvas covers [-size..+size]
		//number of animation frames
		delay  = 8  //delay between frame in 10ms units
	)
	freq := rand.Float64() * 3.0 //relative frequency of y oscillator
	//gif.GIF也是一个复合声明，生成的是一个struct结构体
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0,0,2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t<cycles*2*math.Pi; t+=res {
			x := math.Sin(t)
			y := math.Sin(t*freq+phase)
			img.SetColorIndex(size+int(x*size+0.5),size+int(y*size+0.5),blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay,delay)
		anim.Image = append(anim.Image,img)
	}
	gif.EncodeAll(out,&anim)
}

func main() {
	//The sequence of images is deterministic unless we seed
	//the pseudo-random number generator using the current time.
	//Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

