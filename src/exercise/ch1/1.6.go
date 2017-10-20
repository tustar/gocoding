package main

import (
	"image/color"
	"math/rand"
	"time"
	"os"
	"image/gif"
	"image"
	"math"
	"io"
)

var palette = []color.Color{
	color.RGBA{255, 255, 255, math.MaxUint8},
	color.Black,
	color.RGBA{0, 0, 128, math.MaxUint8},
	color.RGBA{30, 144, 255, math.MaxUint8},
	color.RGBA{72, 61, 139, math.MaxUint8},
	color.RGBA{255, 0, 255, math.MaxUint8},
	color.RGBA{124, 252, 0, math.MaxUint8}}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
	greenIndex = 2 // next color in palette 定义新的颜色
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}
func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i, c := 0, 1; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(c))
		}
		c++
		if c >= len(palette) {
			c = 0
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
