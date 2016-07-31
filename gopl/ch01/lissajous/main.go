// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

//!-main
// Packages not needed by version in book.
import (
	"log"
	"net/http"
	"net/url"
	"time"
	"strconv"
)

//!+main

var palette = []color.Color{
	color.Black,
	color.RGBA{255, 0, 0, 255},
	color.RGBA{0, 255, 0, 255},
	color.RGBA{0, 0, 255, 255},
}

const (
	blackIndex = 0
	redIndex = 1
	greenIndex = 2
	blueIndex = 3
)

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	//!+http
	http.HandleFunc("/", lissajousHandler)
	//!-http
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("URL: ", r.URL.Path)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	switch r.URL.Path {
	case "/red":
		lissajous(w, r.Form, redIndex)

	case "/blue":
		lissajous(w, r.Form, blueIndex)

	default:
		lissajous(w, r.Form, greenIndex)
	}
}

func lissajous(out io.Writer, values url.Values, colorIndex uint8) {
	delay := 8
	if values["delay"] != nil {
		i, err := strconv.Atoi(values["delay"][0])
		if err != nil {
			log.Print(err)
		}
		delay = i
	}
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
