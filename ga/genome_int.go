package ga

import (
	"math/rand"
)

const GenomeIntMaxValue = 1e9

type GenomeInt struct {
	Gene []uint32
}

func NewGenomeInt(length uint32) *GenomeInt {
	var g []uint32 = make([]uint32, length)
	return &GenomeInt{g}
}

func (this *GenomeInt) Randomize() {
	for n, _ := range this.Gene {
		this.Gene[n] = uint32(rand.Int31n(GenomeIntMaxValue))
	}
}

func (this *GenomeInt) Len() int {
	return len(this.Gene)
}
