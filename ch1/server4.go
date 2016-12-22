package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		rand.Seed(time.Now().UTC().UnixNano())
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		i, err := strconv.Atoi(r.Form.Get("cycles"))
		if err != nil {
			i = 5
		}
		lissajous(w, i)
	}) // 個々のリクエストに対してhandler が呼ばれる
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles int) {
	const (
		// cycles  = int(cycles) //発信機xが完了する周回の回数
		res     = 0.001 //回転の分解能
		size    = 100   // 画像キャンバスは[-size..+size]
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0 //発信機yの相対周波数
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

//handler は、リクエストされたURL rのPath要素を返します。
// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
// }
