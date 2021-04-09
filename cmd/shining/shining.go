package main

import (
	"fmt"
	"github.com/ajstarks/gensvg"
	"os"
)

func main() {
	xp := []float64{50, 70, 70, 50, 30, 30}
	yp := []float64{40, 50, 75, 85, 75, 50}
	xl := []float64{0, 0, 50, 100, 100}
	yl := []float64{100, 40, 10, 40, 100}
	bgcolor := "rgb(227,78,25)"
	bkcolor := "rgb(153,29,40)"
	stcolor := "rgb(65,52,44)"
	stwidth := 12.0
	stylefmt := "stroke:%s;stroke-width:%.2f;fill:%s"
	canvas := gensvg.New(os.Stdout)
	width, height := 500.0, 500.0
	canvas.Start(width, height)
	canvas.Def()
	canvas.Gid("unit")
	canvas.Polyline(xl, yl, "fill:none")
	canvas.Polygon(xp, yp)
	canvas.Gend()
	canvas.Gid("runit")
	canvas.TranslateRotate(150, 180, 180)
	canvas.Use(0, 0, "#unit")
	canvas.Gend()
	canvas.Gend()
	canvas.DefEnd()
	canvas.Rect(0, 0, width, height, "fill:"+bgcolor)
	canvas.Gstyle(fmt.Sprintf(stylefmt, stcolor, stwidth, bkcolor))
	for y := 0.0; y < height; y += 130 {
		for x := -50.0; x < width; x += 100 {
			canvas.Use(x, y, "#unit")
			canvas.Use(x, y, "#runit")
		}
	}
	canvas.Gend()
	canvas.End()
}
