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
	return &GenomeInt{Gene: g}
}

func (this *GenomeInt) Randomize() {
	for n, _ := range this.Gene {
		this.Gene[n] = uint32(rand.Int31n(GenomeIntMaxValue))
	}
}

func (this *GenomeInt) Len() uint32 {
	return uint32(len(this.Gene))
}

func (this *GenomeInt) Crossover(g Genome, index uint32) Genome {
	genes := make([]uint32, this.Len())
	copy(genes, this.Gene)
	genes1 := genes[:index]
	genes2 := g.(*GenomeInt).Gene[index:]
	newGenome := NewGenomeInt(this.Len())
	newGenome.Gene = append(genes1, genes2...)
	//newGenome.Gene = genes
	return newGenome
}
