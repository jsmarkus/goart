package main

import (
	"github.com/jsmarkus/goart/graphics"
)

func main() {
	var a = graphics.CreateRaster(50, 10)
	a.SetColor(graphics.Color{0xFF, 0xFF, 0xFF})
	//a.DrawPoint(1, 1)
	//a.DrawPoint(5, 5)
	a.SetColor(graphics.Color{0xAA, 0xAA, 0xAA})
	//a.DrawPoint(0, 0)
	// fmt.Printf("%#v", a)
	//for i := 0; i < 1000000; i++ {
	a.DrawTriangle(25, 0, 49, 3, 0, 9)
	//}
	graphics.PrintRaster(a)
}
