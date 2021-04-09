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
	canvas.Rect(0, 0, width, height, "fill:black")
	for i := 0.0; i < width; i++ {
		x := rand.Float64()*width
		y := rand.Float64()*height
		canvas.Line(x, y, x, y+1, "stroke:white")
	}
	earth := 4.0
	sun := earth * 109
	canvas.Circle(150, 50, earth, "fill:blue")
	canvas.Circle(width, height, sun, "fill:rgb(255, 248, 231)")
	canvas.End()
}
