package analytics

import "github.com/jsmarkus/goart/graphics"

func Asymmetry(raster *graphics.Raster) uint32 {
	var x0, x1, y, sum uint32
	for y = 0; y < raster.Height; y++ {
		for x0 = 0; x0 < raster.Width/2; x0++ {
			x1 = raster.Width - x0 - 1
			c1 := raster.GetColorAt(x0, y)
			c2 := raster.GetColorAt(x1, y)
			sum += ColorDistance(&c1, &c2)
		}
	}
	return sum
}
