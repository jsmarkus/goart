package raster

import (
	. "github.com/jsmarkus/goart/color"
)

type Raster struct {
	buffer []Color
	color  Color
	Width  uint32
	Height uint32
}

func CreateRaster(w uint32, h uint32) *Raster {
	size := w * h
	return &Raster{make([]Color, size), Color{}, w, h}
}

func (r *Raster) SetColor(c Color) {
	r.color = c
}

func (r *Raster) GetColorAt(x, y uint32) Color {
	offset := r.xyToOffset(x, y)
	return r.buffer[offset]
}

func (r *Raster) xyToOffset(x, y uint32) uint32 {
	return x + r.Width*y
}

func (r *Raster) DrawPoint(x, y uint32) {
	offset := r.xyToOffset(x, y)
	r.buffer[offset] = r.color
}

func abs(v int32) uint32 {
	if v < 0 {
		return uint32(-v)
	} else {
		return uint32(v)
	}
}

func sign(v int32) int8 {
	if v < 0 {
		return -1
	} else {
		return 1
	}
}

func (r *Raster) DrawLine(x1, y1, x2, y2 uint32) {
	var deltaX int32 = int32(abs(int32(x2 - x1)))
	var deltaY int32 = int32(abs(int32(y2 - y1)))
	var signX int8 = sign(int32(x2 - x1))
	var signY int8 = sign(int32(y2 - y1))
	var error int32 = deltaX - deltaY
	r.DrawPoint(x2, y2)

	for x1 != x2 || y1 != y2 {
		r.DrawPoint(x1, y1)
		var error2 int32 = error * 2
		if error2 > -deltaY {
			error -= deltaY
			x1 += uint32(signX)
		}
		if error2 < deltaX {
			error += deltaX
			y1 += uint32(signY)
		}
	}
}

func (r *Raster) DrawTriangle(x1, y1, x2, y2, x3, y3 uint32) {
	////http://www.gamedev.ru/code/forum/?id=95311
	if y1 > y2 {
		y1, y2 = y2, y1
		x1, x2 = x2, x1
	}
	if y2 > y3 {
		y3, y2 = y2, y3
		x3, x2 = x2, x3
	}
	if y1 > y2 {
		y1, y2 = y2, y1
		x1, x2 = x2, x1
	}

	if y1 == y2 && y2 == y3 {
		return
	}

	var xb, xe uint32

	for scanline := y1; scanline < y2; scanline++ {
		if x2 > x1 {
			xb = x2 - (y2-scanline)*(x2-x1)/(y2-y1)
		} else {
			xb = x2 + (y2-scanline)*(x1-x2)/(y2-y1)
		}

		if x3 > x1 {
			xe = x3 - (y3-scanline)*(x3-x1)/(y3-y1)
		} else {
			xe = x3 + (y3-scanline)*(x1-x3)/(y3-y1)
		}

		if xb > xe {
			xb, xe = xe, xb
		}

		r.DrawLine(xb, scanline, xe, scanline)
	}

	for scanline := y2; scanline < y3; scanline++ {
		if x2 > x3 {
			xb = x3 + (y3-scanline)*(x2-x3)/(y3-y2)
		} else {
			xb = x3 - (y3-scanline)*(x3-x2)/(y3-y2)
		}

		if x1 > x3 {
			xe = x3 + (y3-scanline)*(x1-x3)/(y3-y1)
		} else {
			xe = x3 - (y3-scanline)*(x3-x1)/(y3-y1)
		}

		if xb > xe {
			xb, xe = xe, xb
		}

		r.DrawLine(xb, scanline, xe, scanline)
	}
}
