package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
	"sync"
)

const (
	backroundIndex  = 0 // first color in pallete
	foregroundIndex = 1 // next color in pallete
)

var palette = []color.Color{color.RGBA{40, 105, 120, 10}, color.White}
var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/header", header)
	http.HandleFunc("/lissajous", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		if relativeFreq, err := strconv.Atoi(r.FormValue("relativeFreq")); err != nil {
			fmt.Fprintf(w, "Form value of \"%q\" must be convertible to float64 type", r.FormValue("relativeFreq"))
		} else {
			lissajous(w, float64(relativeFreq))
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func header(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s, %s, %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func lissajous(out io.Writer, relativeFreq float64) {
	const (
		cycles  = 4   // number of complete x oscillator revolutions
		res     = .01 // angular resolution
		size    = 100 // image canvas covers [-size..+size]
		nframes = 256 //number of animation frams
		delay   = 8   // delay between frames in 10ms units
		// relativeFreq = 1.0 // relative frequency of y oscillator
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
