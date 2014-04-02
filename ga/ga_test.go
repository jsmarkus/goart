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

func rateDifference(gn1, gn2 *GenomeInt) (diffQuantity, maxDeviation int) {
	maxDeviation = 0
	diffQuantity = 0
	for i, g := range gn1.Gene {
		dev := int(g) - int(gn2.Gene[i])
		if dev < 0 {
			dev = -dev
		}
		if dev != 0 {
			diffQuantity++
		}
		if maxDeviation < int(dev) {
			maxDeviation = int(dev)
		}
	}
	return
}

func TestMutate(t *testing.T) {
	g := NewGenomeInt(100000)
	g.Randomize()
	g2 := g.Copy().(*GenomeInt)
	g.Mutate(1e-7, 0.02)

	diffQuantity, maxDeviation := rateDifference(g, g2)
	fmt.Printf("%d : %d\n", diffQuantity, maxDeviation)
	if diffQuantity < 1900 || diffQuantity > 2100 {
		t.Errorf("Quantity of mutated genes (%d) differs from the expected (1900-2100)", diffQuantity)
	}
	if maxDeviation < 90 || maxDeviation > 100 {
		t.Errorf("Max deviation value (%d) differs from the expected (90-100)", maxDeviation)
	}
}
