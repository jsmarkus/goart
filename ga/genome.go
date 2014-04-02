package ga

type Genome interface {
	Randomize()
	Len() uint32
	Crossover(g Genome, index uint32) Genome
	Mutate(strength, probability float32)
}
