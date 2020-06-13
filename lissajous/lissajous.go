package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

//[]color.Color{...} is a composite literal that can be used to instantiate any of Go's composite types
//from a sequence of element values
//palette here is a slice
//var palette = []color.Color{color.White, color.Black} //color.Color belongs to the image/color package
//ex1.5 change the color palette to green on black
var palette = []color.Color{color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xFF}}

//const can be package level or function level and can have values of type number, string or boolean
const (
	foregroundIndex = 0 //first color in palette
	backgroundIndex = 1 //next color in palette
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     //number of complete x oscillator revolutions
		res     = 0.001 //angular resolution
		size    = 100   //image canvas covers [-size..+size]
		nframes = 64    //number of animation frames
		delay   = 8
	)

	freq := rand.Float64() * 3.0 //relative frequency of y oscillator
	//gif.GIF{...} is another example of composite literal that instantiates a struct
	//a struct is a group of values called fields, often of different types, that are collected together
	//in a single object that can be treated as a unit.
	anim := gif.GIF{LoopCount: nframes} //gif.GIF belongs to image/gif package
	phase := 0.0                        //phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1) //image.Rect belongs to image package
		img := image.NewPaletted(rect, palette)      //image.NewPaletted belongs to image package
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), backgroundIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
