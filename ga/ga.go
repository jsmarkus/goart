package ga

import ()

type Ga struct {
	Fitness          Fitness
	GenerationLength uint32
	GenomeLength     uint32
}

type Fitness func(gnm Genome)

func NewGa(generationLength uint32, genomeLength uint32, fitness Fitness) *Ga {
	a := &Ga{
		GenerationLength: generationLength,
		GenomeLength:     genomeLength,
		Fitness:          fitness,
	}
	return a
}

func (this *Ga) Step() {
	gnr := this.Breed()
	gnr.Randomize()
	genomes := this.Select(gnr, 10) //TODO: how many

}

func (this *Ga) Select(gnr *Generation, number uint32) []Genome {
}

func (this *Ga) Breed() *Generation {
	gnr := NewGeneration(
		this.GenerationLength, this.GenomeLength, NewGenomeIntAsGenome) //TODO: make constructor (NewGenomeIntAsGenome) configurable
	return gnr
}
