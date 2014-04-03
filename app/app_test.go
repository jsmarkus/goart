package app

import (
	//"fmt"
	"github.com/jsmarkus/goart/ga"
	"github.com/jsmarkus/goart/graphics"
	"testing"
)

func TestTriangleRenderer(t *testing.T) {
	gnm := ga.NewGenomeInt(9 * 15)
	rst := graphics.CreateRaster(50, 10)

	tr := NewTriangleRenderer(gnm, rst)
	gnm.Randomize()
	tr.Render()
	//fmt.Println(tr.Raster)
	graphics.PrintRaster(tr.Raster)
}
