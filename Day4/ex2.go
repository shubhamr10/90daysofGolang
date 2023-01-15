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

// lissajous generates GIF animation of random of Lissajous figures

var palette3 = []color.Color{color.White, color.Black, color.RGBA{
	R: 12,
	G: 34,
	B: 56,
	A: 7,
}, color.RGBA{
	R: 112,
	G: 34,
	B: 56,
	A: 7,
}, color.RGBA{
	R: 12,
	G: 34,
	B: 6,
	A: 7,
}}

const (
	whiteIndex2 = 0 // first color in palette
	blackIndex2 = 1 // second color in palette
)

func main() {
	lissajous3(os.Stdout)
}

func lissajous3(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas cover [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette3)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(5 - 0 + 1)
			un := uint8(n)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), un)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
