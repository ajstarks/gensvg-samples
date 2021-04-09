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
	tiles, maxstroke := 25.0, 10.0
	rand.Seed(time.Now().Unix())
	canvas.Start(width, height)
	linecaps := []string{"butt", "round", "square"}
	strokefmt := "stroke-width:%.2f"
	lcfmt := "stroke:black;stroke-linecap:%s"
	canvas.Gstyle(fmt.Sprintf(lcfmt, linecaps[rand.Intn(3)]))
	var sw string
	for y := 0.0; y < tiles; y++ {
		for x := 0.0; x < tiles; x++ {
			px := width / tiles * x
			py := height / tiles * y
			if rand.Intn(100) > 50 {
				sw = fmt.Sprintf(strokefmt, rand.Float64()*(maxstroke)+1)
				canvas.Line(px, py, px+width/tiles, py+height/tiles, sw)
			} else {
				sw = fmt.Sprintf(strokefmt, rand.Float64()*(maxstroke)+1)
				canvas.Line(px, py+height/tiles, px+width/tiles, py, sw)
			}
		}
	}
	canvas.Gend()
	canvas.End()
}
