package analytics

import (
	"github.com/jsmarkus/goart/graphics"
	. "github.com/jsmarkus/goart/math"
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
	return ISqrt(r*r + g*g + b*b)
}
