package main

import (
	"fmt"
	"github.com/jsmarkus/goart/analytics"
	"github.com/jsmarkus/goart/app"
	"github.com/jsmarkus/goart/ga"
	"github.com/jsmarkus/goart/graphics"
)

const NG = 100000

func experiment(pipe chan string) {
	var generation [NG]*ga.GenomeInt
	for i := 0; i < NG; i++ {
		genome := ga.NewGenomeInt(5 * 9)
		genome.Randomize()
		generation[i] = genome
	}

	raster := graphics.CreateRaster(10, 10)
	renderer := app.NewTriangleRenderer(raster)

	for i := 0; i < NG; i++ {
		renderer.Render(generation[i])
		as := analytics.Asymmetry(raster)
		if as <= 1000 {
			pipe <- fmt.Sprintf("%d: as=%d\n", i, as)
			//fmt.Printf("%d: as=%d\n", i, as)
			//graphics.PrintRaster(raster)
		}
	}

	close(pipe)

}

func main() {
	pipe := make(chan string)
	go experiment(pipe)
	for msg := range pipe {
		fmt.Println(msg)
	}
}
