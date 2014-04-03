package analytics

import (
	"github.com/jsmarkus/goart/graphics"
	"testing"
)

func TestSymmetry(t *testing.T) {
	rst := graphics.CreateRaster(4, 4)
	Symmetry(rst)
}
