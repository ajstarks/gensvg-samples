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

// inspired by Gerhard Richter's 256 colors, 1974
func main() {
	rand.Seed(time.Now().Unix())
	canvas.Decimals=0
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height)

	w, h, gutter := 24.0, 18.0, 5.0
	rows, cols := 16.0, 16.0
	top, left := 20.0, 20.0

	for r, x := 0.0, left; r < rows; r++ {
		for c, y := 0.0, top; c < cols; c++ {
			canvas.Rect(x, y, w, h,
				canvas.RGB(rand.Intn(255), rand.Intn(255), rand.Intn(255)))
			y += (h + gutter)
		}
		x += (w + gutter)
	}
	canvas.End()
}
