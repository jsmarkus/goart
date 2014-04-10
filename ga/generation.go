package ga

import (
	"math/rand"
)

type Generation interface {
}

type GenerationS struct {
	Genome []Genome
}

func NewGeneration(length uint32, genomeLength uint32, constr GenomeConstructor) *GenerationS {
	gnm := make([]Genome, length)
	for i := uint32(0); i < length; i++ {
		gnm[i] = constr(genomeLength)
	}
	return &GenerationS{
		Genome: gnm,
	}
}

func (this *GenerationS) Mutate(strength, geneProbability, genomeProbability float32) {
	for _, gnm := range this.Genome {
		if rand.Float32() < genomeProbability {
			gnm.Mutate(strength, geneProbability)
		}
	}
}
func (this *GenerationS) Randomize() {
	for _, gnm := range this.Genome {
		gnm.Randomize()
	}
}
