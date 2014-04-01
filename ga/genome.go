package ga

type Genome interface {
	Randomize()
	Len() int
}
