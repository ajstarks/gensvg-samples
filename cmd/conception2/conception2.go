package main

import (
	"github.com/ajstarks/gensvg"
	"math/rand"
	"os"
	"time"
)

var (
	canvas = gensvg.New(os.Stdout)
	width  = 500.0
	height = 500.0
)

func male(x, y, w float64) {
	canvas.Ellipse(x, y, w, w/2, "fill:blue")
	canvas.Bezier(
		x-(w*8), y,
		x-(w*4), y-(w*4),
		x-(w*4), y+w,
		x-w, y, "stroke:blue;fill:none")
}

func female(x, y, w float64) {
	canvas.Circle(x, y, w, "fill:pink")
}

func main() {
	rand.Seed(time.Now().Unix())
	msize := 5.0
	fsize := msize * 40
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, "fill:white")
	female(width, height-50, fsize)
	for i := 0; i < 100; i++ {
		canvas.TranslateRotate((rand.Float64()*300)+100, (rand.Float64()*200)+200, rand.Float64()*45)
		male(0, 0, msize)
		canvas.Gend()
	}
	canvas.End()
}
