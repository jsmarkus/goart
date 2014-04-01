// genome.go
package ga

type Genome interface {
	Randomize()
	Len() int
}

type GenomeInt struct {
	Gene []uint32
}

func NewGenomeInt(length uint32) *GenomeInt {
	var g []uint32 = make([]uint32, length)
	return &GenomeInt{g}
}

func (GenomeInt) Randomize() {
	//TODO
}

func (this GenomeInt) Len() int {
	return len(this.Gene)
}
