//Server is a minimal echo server and counter server
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex

//count keeps track of count of all requests to our server
var count int
var cycles float64 = 5 //number of complete x oscillator revolutions

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xFF}}

//const can be package level or function level and can have values of type number, string or boolean
const (
	foregroundIndex = 0 //first color in palette
	backgroundIndex = 1 //next color in palette
)

func main() {
	//there are two handlers, request determines which one is called
	http.HandleFunc("/", handler) //each request calls handler
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", lissajous)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

//A request is represented by a struct of type http.Request, URL is one of the fields in the struct
func handler(w http.ResponseWriter, r *http.Request) {
	//if two concurrent requests try to update count at the same time, it might not be incremented consistently
	//mu.Lock() and mu.Unlock() is used to avoid this specific race condition when updating count
	//mu.Lock() and mu.Unlock() cradle each access to count
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	//echo http request
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q]= %q\n", k, v)
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

//counter echoes the number of calls so far
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

//lissajous handler
func lissajous(w http.ResponseWriter, r *http.Request) {
	const (
		//cycles  = 5     //number of complete x oscillator revolutions
		res     = 0.001 //angular resolution
		size    = 100   //image canvas covers [-size..+size]
		nframes = 64    //number of animation frames
		delay   = 8
	)

	//ex1.12 extracts cycles from http://localhost:8080/lissajous?cycles=500 and sets the value of the package variable cycles instead of setting it as a default const
	for k, v := range r.Form {
		fmt.Println(k, v)
		if k == "cycles" {
			cyclec, err := strconv.Atoi(k)
			if err != nil {
				log.Print(err)
			}
			cycles = float64(cyclec)
		}
	}

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
	gif.EncodeAll(w, &anim)
}
