package main

import (
	"github.com/ajstarks/gensvg"
	"math/rand"
	"os"
)

func tloc(x, y, s float64, r, d float64) (float64, float64) {
	padding := 2 * s
	return (padding + (x * s) - (.5 * s) + (r * d)),
		(padding + (y * s) - (.5 * s) + (r * d))
}

func random(n float64) float64 {
	x := rand.Float64()
	if x < 0.5 {
		return -n * x
	}
	return n * x
}

func main() {
	columns, rows, sqrsize := 12.0, 12.0, 32.0
	rndStep, dampen := .22, 0.45
	width, height := (columns+4)*sqrsize, (rows+4)*sqrsize
	canvas := gensvg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, "fill:white")
	canvas.Gstyle("fill:rgb(0,0,127);fill-opacity:0.3")
	for y, randsum := 1.0, 0.0; y <= rows; y++ {
		randsum += float64(y) * rndStep
		for x := 1.0; x <= columns; x++ {
			tx, ty := tloc(x, y, sqrsize, random(randsum), dampen)
			canvas.TranslateRotate(tx, ty, random(randsum))
			canvas.CenterRect(0, 0, sqrsize, sqrsize)
			canvas.Gend()
		}
	}
	canvas.Gend()
	canvas.End()
}
