package analytics

import (
	"github.com/jsmarkus/goart/graphics"
	"testing"
)

func TestAsymmetry(t *testing.T) {
	rst := graphics.CreateRaster(4, 4)
	rst.SetColor(graphics.Color{5, 60, 71})
	if Asymmetry(rst) > 0 {
		t.Error("assymetry is wrong")
	}
	rst.DrawPoint(2, 1)
	if Asymmetry(rst) <= 0 {
		t.Error("assymetry is wrong")
	}
	rst.DrawPoint(1, 1)
	if Asymmetry(rst) > 0 {
		t.Error("assymetry is wrong")
	}
}

func TestColorDistance(t *testing.T) {
	if ColorDistance(
		&graphics.Color{15, 15, 15},
		&graphics.Color{15, 15, 15},
	) != 0 {
		t.Error("color distance is wrong")
	}

	if ColorDistance(
		&graphics.Color{15, 15, 15},
		&graphics.Color{15, 15, 16},
	) <= 0 {
		t.Error("color distance is wrong")
	}
}
