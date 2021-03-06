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

func coord(x, y, size float64, label string) {
	tstyle := "text-anchor:middle;font-size:14pt"
	offset := size + (size / 2)
	canvas.Text(x, y-offset,
		fmt.Sprintf("%s (%d,%d)", label, int(x), int(y)), tstyle)
	canvas.Circle(x, y, size)
}

func showcurve(bx, by, cx, cy, ex, ey float64) {
	dotsize := 5.0
	sw := dotsize * 2
	cfmt := "stroke:%s;stroke-width:%.2f;fill:none;stroke-opacity:%.2f"
	style := fmt.Sprintf(cfmt, "red", sw, 0.2)
	coord(bx, by, dotsize, "begin")
	coord(ex, ey, dotsize, "end")
	coord(cx, cy, dotsize, "control")
	canvas.Qbez(bx, by, cx, cy, ex, ey, style)
}

func main() {
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, "fill:rgb(250,250,250)")
	canvas.Grid(0, 0, width, height, 25, "stroke:lightgray")
	showcurve(70, 200, 100, 425, 425, 125)
	canvas.End()
}
