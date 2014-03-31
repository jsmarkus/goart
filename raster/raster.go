package raster
import . "github.com/jsmarkus/goart/color"

type Raster struct {
    buffer []Color
    color Color
    Width uint32
    Height uint32
}

func CreateRaster(w uint32, h uint32) *Raster {
    size := w * h
    return &Raster{make([]Color, size), Color{}, w, h}
}

func (r *Raster) SetColor (c Color) {
    r.color = c
}

func (r *Raster) DrawPoint (x, y uint32) {
    offset := r.xyToOffset(x, y)
    r.buffer[offset] = r.color
}

func (r *Raster) GetColorAt (x, y uint32) Color {
    offset := r.xyToOffset(x, y)
    return r.buffer[offset]
}

func (r *Raster) xyToOffset (x, y uint32) uint32 {
    return x + r.Height * y
}