package main

import (
	"math/rand"
	"os"

	"github.com/ajstarks/gensvg"
)

var (
	canvas = gensvg.New(os.Stdout)
	width  = 500.0
	height = 500.0
)

func main() {
	canvas.Start(width, height)
	canvas.Gstyle("fill:white")
	var color string
	radius := 80.0
	step := 8.0
	for i := 0; i < 200; i++ {
		if i%4 == 0 {
			color = "rgb(127,0,0)"
		} else {
			color = "rgb(0,127,0)"
		}
		x, y := rand.Float64()*(width), rand.Float64()*(height)
		for r, nc := radius, 0; nc < 10; nc++ {
			canvas.Circle(x, y, r, "stroke:"+color)
			r -= step
		}
	}
	canvas.Gend()
	canvas.End()
}
