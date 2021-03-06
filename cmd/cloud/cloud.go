package main

import (
	"os"

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
	canvas.Start(width, height)
	cloud(width/2, height/2, 100, canvas.RGB(127, 127, 127))
	canvas.End()
}
