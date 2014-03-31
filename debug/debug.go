package debug

import (
    "fmt"
  . "github.com/jsmarkus/goart/raster"
)

func PrintRaster(r *Raster) {
    var x, y, c uint32
    for y = 0; y < r.Height; y++ {
      for x = 0; x < r.Width; x++ {
        color := r.GetColorAt(x, y)
        c = color.Serialize()
        if c > 0 {
          fmt.Print("*")
        } else {
          fmt.Print(".")
        }
      }
      fmt.Print("\n")
    }
}