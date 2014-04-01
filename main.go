package main

import (
	"fmt"
	"github.com/jsmarkus/goart/ga"
	//"github.com/jsmarkus/goart/graphics"
)

func main() {
	//var a = graphics.CreateRaster(50, 10)

	var gnm *ga.GenomeInt = ga.NewGenomeInt(10)
	fmt.Printf("%#v\n", gnm)
	fmt.Printf("%#v\n", gnm.Len())
	gnm.Randomize()
	fmt.Printf("%#v\n", gnm)
}
