package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"os"
)

var palette = []color.Color{color.RGBA{40, 105, 120, 10}, color.White}

const (
	backroundIndex  = 0 // first color in pallete
	foregroundIndex = 1 // next color in pallete
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles       = 4   // number of complete x oscillator revolutions
		res          = .01 // angular resolution
		size         = 100 // image canvas covers [-size..+size]
		nframes      = 256 //number of animation frams
		delay        = 8   // delay between frames in 10ms units
		relativeFreq = 1.0 // relative frequency of y oscillator
	)

	anim := gif.GIF{LoopCount: nframes}
	phase := 0.01 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*relativeFreq + phase)
			img.SetColorIndex(size+int(x*size), size+int(y*size), foregroundIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Note: ignoring encoding errors
}
