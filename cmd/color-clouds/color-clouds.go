package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/ajstarks/gensvg"
)

var (
	canvas = gensvg.New(os.Stdout)
	width  = 500.0
	height = 500.0
)

func cloud(x, y, r float64, style string) {
	small := r / 2
	medium := (r * 6) / 10
	canvas.Gstyle(style)
	canvas.Circle(x, y, r)
	canvas.Circle(x+r, y+small, small)
	canvas.Circle(x-r-small, y+small, small)
	canvas.Circle(x-r, y, medium)
	canvas.Rect(x-r-small, y, r*2+small, r)
	canvas.Gend()
}

func main() {
	rand.Seed(time.Now().Unix())
	canvas.Start(width, height)
	for i := 0; i < 50; i++ {
		red := rand.Intn(255)
		green := rand.Intn(255)
		blue := rand.Intn(255)
		size := rand.Float64()*60
		x := rand.Float64()*width
		y := rand.Float64()*height
		cloud(x, y, size, canvas.RGB(red, green, blue))
	}
	canvas.End()
}
