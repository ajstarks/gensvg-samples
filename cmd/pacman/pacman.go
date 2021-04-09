package main

import (
	"github.com/ajstarks/gensvg"
	"os"
)

var (
	canvas = gensvg.New(os.Stdout)
	width  = 500.0
	height = 500.0
)

func main() {
	angle, cx, cy := 30.0, width/2, height/2
	r := width / 4
	p := r / 8

	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height)
	canvas.Gstyle("fill:white")
	for x := 0.0; x < 100; x += 12 {
		if x < 50 {
			canvas.Circle((width*x)/100, cy, p, "fill-opacity:0.5")
		} else {
			canvas.Circle((width*x)/100, cy, p, "fill-opacity:1")
		}
	}
	canvas.Gend()

	canvas.Gstyle("fill:yellow")
	canvas.TranslateRotate(cx, cy, -angle)
	canvas.Arc(-r, 0, r, r, 30, false, true, r, 0)
	canvas.Gend()

	canvas.TranslateRotate(cx, cy, angle)
	canvas.Arc(-r, 0, r, r, 30, false, false, r, 0)
	canvas.Gend()

	canvas.Gend()
	canvas.End()
}
