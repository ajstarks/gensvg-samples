package main

import (
	"math"
	"os"
	"github.com/ajstarks/gensvg"
)
var canvas, width, height = gensvg.New(os.Stdout), 500.0, 500.0

func polar(cx, cy, r, t float64) (float64, float64) {
	return cx + (r * math.Cos(t)), cy + (r * math.Sin(t))
}

func star(x, y float64, n int, inner, outer float64, style string) {
	t, xv, yv := math.Pi/float64(n), make([]float64, n*2), make([]float64, n*2)
	for i := 0; i < n*2; i++ {
		if i%2 == 0 {
			xv[i], yv[i] = polar(0, 0, outer, t*float64(i))
		} else {
			xv[i], yv[i] = polar(0, 0, inner, t*float64(i))
		}
	}
	canvas.TranslateRotate(x, y, 54)
	canvas.Polygon(xv, yv, style)
	canvas.Gend()
}

func aline(x, y, r, a1, a2 float64) {
	x1, y1 := polar(x, y, r, a1)
	x2, y2 := polar(x, y, r, a2)
	canvas.Line(x1, y1, x2, y2, "stroke:maroon;stroke-width:10")
}

func main() {
	x, y, p4, r := width/2, height/2, math.Pi/4, 65.0
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, canvas.RGB(240, 240, 240))
	canvas.Circle(x, y, width/2, canvas.RGB(255, 255, 255))
	star(x, y, 5, 90, 240, canvas.RGB(200, 200, 200))
	aline(x, y, r, p4, 5*p4)
	aline(x, y, r, 3*p4, 7*p4)
	canvas.End()
}
