package analytics

import (
	"github.com/jsmarkus/goart/graphics"
)

func uAbsDif(a, b uint8) uint8 {
	if a > b {
		return a - b
	}
	return b - a
}

func ColorDistance(c1, c2 *graphics.Color) uint32 {
	r := uint32(uAbsDif(c1.R, c2.R))
	g := uint32(uAbsDif(c1.G, c2.G))
	b := uint32(uAbsDif(c1.B, c2.B))
	return r*r + g*g + b*b //TODO: Sqrt()? ha-ha-ha!!!
}

func Asymmetry(raster *graphics.Raster) uint32 {
	//TODO: implement and use color equality testing function
	var x, y0, y1, sum uint32
	for x = 0; x < raster.Width; x++ {
		for y0 = 0; y0 < raster.Height/2; y0++ {
			y1 = raster.Height - y0 - 1
			c1 := raster.GetColorAt(x, y0)
			c2 := raster.GetColorAt(x, y1)
			sum += ColorDistance(&c1, &c2)
		}
	}
	return sum
}
