package main

import (
	"fmt"
	"github.com/jsmarkus/goart/analytics"
	"github.com/jsmarkus/goart/app"
	"github.com/jsmarkus/goart/ga"
	"github.com/jsmarkus/goart/graphics"
)

const NG = 1000000

func main() {
	var generation [NG]*ga.GenomeInt
	for i := 0; i < NG; i++ {
		genome := ga.NewGenomeInt(5 * 9)
		genome.Randomize()
		generation[i] = genome
	}

	for i := 0; i < NG; i++ {
		renderer := app.NewTriangleRenderer(
			generation[i],
			graphics.CreateRaster(10, 10))
		renderer.Render()
		as := analytics.Asymmetry(renderer.Raster)
		if as <= 600 {
			fmt.Printf("%d: as=%d\n", i, as)
			graphics.PrintRaster(renderer.Raster)
		}
	}
}
