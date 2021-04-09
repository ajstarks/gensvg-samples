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

func main() {
	fonts := []string{
		"Helvetica", "Times", "Courier",
		"sans-serif", "serif", "monospace",
	}
	sizes := []float64{10, 12, 16, 21, 24, 36, 48}
	largest := sizes[len(sizes)-1]
	gutter := largest + (largest / 3)
	margin := gutter * 2
	y := 100.0

	canvas.Start(width, height)
	for _, f := range fonts {
		x := margin
		canvas.Gstyle("font-family:" + f)
		canvas.Text(x-10, y, f, "text-anchor:end")
		for _, s := range sizes {
		canvas.Text(x, y, fmt.Sprintf("%0.f", s), fmt.Sprintf("font-size:%0.fpt", s))
			x += s * 2
		}
		canvas.Gend()
		y += gutter
	}
	canvas.End()
}
