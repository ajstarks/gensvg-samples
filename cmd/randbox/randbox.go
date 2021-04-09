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

func main() {
	canvas.Start(width, height)
	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i++ {
		fill := canvas.RGBA(
			rand.Intn(255),
			rand.Intn(255),
			rand.Intn(255),
			rand.Float64())
		canvas.Rect(
			rand.Float64()*width,
			rand.Float64()*height,
			rand.Float64()*100,
			rand.Float64()*100,
			fill)
	}
	canvas.End()
}
