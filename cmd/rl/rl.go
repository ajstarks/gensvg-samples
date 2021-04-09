package main

import (
	"fmt"
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

func main() {
	rand.Seed(time.Now().Unix())
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height)
	canvas.Gstyle("stroke-width:10")

	for i := 0.0; i < width; i++ {
		r := rand.Intn(255)
		canvas.Line(i, 0, rand.Float64()*width, height,
			fmt.Sprintf("stroke:rgb(%d,%d,%d); opacity:0.39", r, r, r))
	}
	canvas.Gend()
	canvas.End()
}
