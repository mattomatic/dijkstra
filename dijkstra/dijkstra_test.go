package dijkstra

import (
    "testing"
    "github.com/mattomatic/dijkstra/graph"
)

func TestComputeCut(t *testing.T) {
    g1 := graph.NewGraph()
    g2 := graph.NewGraph()
    
    a := graph.NewNode(1)
    b := graph.NewNode(2)
    e := graph.NewEdge(a, b, 1)
    a.AddEdges(e)
    
    g1.AddNodes(a)
    g2.AddNodes(b)
    
    cut := getCut(g1, g2)
    
    edge := <-cut
    
    if edge != e {
        testing.Error()
    }
}