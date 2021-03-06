package main

import (
	"os"

	"github.com/ajstarks/gensvg"
)

var (
	canvas = gensvg.New(os.Stdout)
	width  = 500.0
	height = 500.0
)

func main() {
	lorem := []string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing",
		"elit, sed do eiusmod tempor incididunt ut labore et",
		"dolore magna aliqua. Ut enim ad minim veniam, quis",
		"nostrud exercitation ullamco laboris nisi ut aliquip",
		"ex ea commodo consequat. Duis aute irure dolor in",
		"reprehenderit in voluptate velit esse cillum dolore eu",
		"fugiat nulla pariatur. Excepteur sint occaecat cupidatat",
		"non proident, sunt in culpa qui officia deserunt mollit",
	}
	fontlist := []string{"Georgia", "Helvetica", "Gill Sans"}
	size, leading := 14.0, 16.0
	x, y := 50.0, 20.0
	tsize := float64(len(lorem))*leading + size*3
	canvas.Start(width, height)
	for _, f := range fontlist {
		canvas.Gstyle("font-family:" + f)
		canvas.Textlines(x, y, lorem, size, leading, "black", "start")
		canvas.Text(x, size+y+tsize/2, f, "fill-opacity:0.3;fill:red;font-size:750%")
		canvas.Gend()
		y += tsize
	}
	canvas.End()
}
