package ga

import (
	"fmt"
	"testing"
)

func TestGenomeInt(t *testing.T) {
	g := NewGenomeInt(10)
	g.Randomize()

	if g.Len() != 10 {
		t.Error("length")
	}
}
func TestCrossover(t *testing.T) {
	g1 := NewGenomeInt(4)
	g1.Gene = []uint32{1, 2, 3, 4}

	g2 := NewGenomeInt(4)
	g2.Gene = []uint32{5, 6, 7, 8}

	g3 := g1.Crossover(g2, 2).(*GenomeInt)
	if fmt.Sprintf("%v", g1.Gene) != "[1 2 3 4]" {
		t.Error("g1")
	}
	if fmt.Sprintf("%v", g2.Gene) != "[5 6 7 8]" {
		t.Error("g2")
	}
	if fmt.Sprintf("%v", g3.Gene) != "[1 2 7 8]" {
		t.Error("g3")
	}

	t.Logf("%#v", g3)
}
