package app

import (
	//"fmt"
	"github.com/jsmarkus/goart/ga"
	"github.com/jsmarkus/goart/graphics"
)

const geneChainLen = 9 //color: 3; coords: 3*2

type TriangleRenderer struct {
	Raster *graphics.Raster
}

func NewTriangleRenderer(raster *graphics.Raster) *TriangleRenderer {
	return &TriangleRenderer{
		Raster: raster,
	}
}

func (this *TriangleRenderer) Render(genome *ga.GenomeInt) {
	this.Raster.Clear()
	var colorDivider = uint32(ga.GenomeIntMaxValue) / 255
	var xDivider = uint32(ga.GenomeIntMaxValue) / this.Raster.Width
	var yDivider = uint32(ga.GenomeIntMaxValue) / this.Raster.Height
	for i := uint32(0); i < genome.Len(); i += geneChainLen {
		offset := i
		color := graphics.Color{
			R: uint8(genome.Gene[offset] / colorDivider),
			G: uint8(genome.Gene[offset+1] / colorDivider),
			B: uint8(genome.Gene[offset+2] / colorDivider),
		}
		//fmt.Println(color)
		x1 := genome.Gene[offset+3] / xDivider
		y1 := genome.Gene[offset+4] / yDivider

		x2 := genome.Gene[offset+5] / xDivider
		y2 := genome.Gene[offset+6] / yDivider

		x3 := genome.Gene[offset+7] / xDivider
		y3 := genome.Gene[offset+8] / yDivider

		//fmt.Printf("(%d,%d) (%d,%d) (%d,%d)\n", x1, y1, x2, y2, x3, y3)
		this.Raster.SetColor(color)
		this.Raster.DrawTriangle(x1, y1, x2, y2, x3, y3)
	}
}
