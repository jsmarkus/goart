package main

import (
	// "fmt"
	. "github.com/jsmarkus/goart/color"
	"github.com/jsmarkus/goart/debug"
	. "github.com/jsmarkus/goart/raster"
)

func main() {
	var a = CreateRaster(10, 10)
	a.SetColor(Color{0xFF, 0xFF, 0xFF})
	//a.DrawPoint(1, 1)
	//a.DrawPoint(5, 5)
	a.SetColor(Color{0xAA, 0xAA, 0xAA})
	//a.DrawPoint(0, 0)
	// fmt.Printf("%#v", a)
	a.DrawLine(1, 1, 5, 5)
	debug.PrintRaster(a)
}
