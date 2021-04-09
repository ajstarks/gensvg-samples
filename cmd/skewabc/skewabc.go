package main

import (
	"fmt"
	"os"

	"github.com/ajstarks/gensvg"
)

var (
	canvas = gensvg.New(os.Stdout)
	width  = 500.0
	height = 500.0
)

func sky(x, y, w, h, a float64, s string) {
	tfmt := "font-family:sans-serif;font-size:%.2fpx;text-anchor:middle"
	canvas.Gstyle(fmt.Sprintf(tfmt, w/2))
	canvas.SkewXY(0, float64(a))
	canvas.Rect(x, y, w, h, "fill:black;fill-opacity:0.3")
	canvas.Text(x+w/2, y+h/2, s, "fill:white;baseline-shift:-33%")
	canvas.Gend()
	canvas.Gend()
}

func main() {
	canvas.Start(width, height)
	canvas.Grid(0, 0, width, height, 50, "stroke:lightblue")
	sky(100, 100, 100, 100, 30, "A")
	sky(200, 332, 100, 100, -30, "B")
	sky(300, -15, 100, 100, 30, "C")
	canvas.End()
}
