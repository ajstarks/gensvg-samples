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

func main() {
	rg := []gensvg.Offcolor{
		{1, "powderblue", 1},
		{10, "lightskyblue", 1},
		{100, "darkblue", 1},
	}
	lg := []gensvg.Offcolor{
		{10, "black", 1},
		{20, "gray", 1},
		{100, "lightgray", 1},
	}
	canvas.Start(width, height)
	canvas.Def()
	canvas.RadialGradient("rg", 50, 50, 50, 30, 30, rg)
	canvas.LinearGradient("lg", 0, 100, 0, 0, lg)
	canvas.DefEnd()
	canvas.Circle(width/2, height-300, 100, "fill:url(#rg)")
	canvas.Ellipse(width-110, height-50, 100, 20, "fill:url(#lg)")
	canvas.End()
}
