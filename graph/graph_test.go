package graph

import (
	"testing"
)

func TestComputeCut(t *testing.T) {
	g1 := NewGraph()
	g2 := NewGraph()

	a := NewNode(1)
	b := NewNode(2)
	e := NewEdge(a, b, 1)
	a.AddEdges(e)

	g1.AddNodes(a)
	g2.AddNodes(b)

	cut := g1.GetCut(g2)

	edge := <-cut

	if edge != e {
		t.Error()
	}
}