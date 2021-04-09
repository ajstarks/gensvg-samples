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
	pct := 5.0
	pw, ph := (width*pct)/100, (height*pct)/100
	canvas.Start(width, height)

	canvas.Def()
	canvas.Pattern("hatch", 0, 0, pw, ph, "user")
	canvas.Gstyle("fill:none;stroke-width:1")
	canvas.Path(fmt.Sprintf("M0,0 l%0.f,%0.f", pw, ph), "stroke:red")
	canvas.Path(fmt.Sprintf("M%0.f,0 l-%0.f,%0.f", pw, pw, ph), "stroke:blue")
	canvas.Gend()
	canvas.PatternEnd()
	canvas.DefEnd()

	x1 := width / 2
	x2 := (width * 4) / 5
	canvas.Gstyle("stroke:black; font-size: 72pt; text-anchor:middle; fill:url(#hatch)")
	canvas.Circle(x1, height/2, height/8)
	canvas.CenterRect(x2, height/2, height/4, height/4)
	canvas.Text(x1, height-50, "Go")
	canvas.Text(x2, height-50, "fast")
	canvas.Gend()
	canvas.End()
}
