package main

import (
	"math/rand"
	"os"
	"time"
	"github.com/ajstarks/gensvg"
)
var (
	canvas = gensvg.New(os.Stdout)
	width  = 500.0
	height = 500.0
)
func randcolor() string {
	return canvas.RGB(rand.Intn(255),rand.Intn(255),rand.Intn(255))
}
func rcube(x, y, l float64) {
	l2, l3, l4, l6, l8 := l*2, l*3, l*4, l*6, l*8
	tx := []float64{x, x + (l3), x, x - (l3), x}
	ty := []float64{y, y + (l2), y + (l4), y + (l2), y}
	lx := []float64{x - (l3), x, x, x - (l3), x - (l3)}
	ly := []float64{y + (l2), y + (l4), y + (l8), y + (l6), y + (l2)}
	rx := []float64{x + (l3), x + (l3), x, x, x + (l3)}
	ry := []float64{y + (l2), y + (l6), y + (l8), y + (l4), y + (l2)}
	canvas.Polygon(tx, ty, randcolor())
	canvas.Polygon(lx, ly, randcolor())
	canvas.Polygon(rx, ry, randcolor())
}
func main() {
	rand.Seed(time.Now().Unix())
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height)
	xp, y := width/10, height/10
	n, hspace, vspace, size := 3, width/5, height/4, width/40
	for r := 0; r < n; r++ {
		for x := xp; x < width; x += hspace {
			rcube(x, y, size)
		}
		y += vspace
	}
	canvas.End()
}
