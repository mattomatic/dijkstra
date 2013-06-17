package dijkstra

import (
    "github.com/mattomatic/dijkstra/graph"
)

type PathMap map[*graph.Node]int
const MAXWEIGHT = 1000000 // insert any biggish number here

// Compute the shortest path from the source node to each other node in the graph
func Dijkstra(g *graph.Graph, s *graph.Node) (PathMap) {
    p := make(PathMap)
    v := graph.NewGraph()
    u := g
    
    initialize(v, u, s, p)

    for u.Size() > 0 {
        step(s, v, u, p)
    }

    return p
}

func initialize(v *graph.Graph, u *graph.Graph, s *graph.Node, p PathMap) {
    v.AddNodes(s)
    u.RemoveNodes(s)
    p[s] = 0
}

func step(s *graph.Node, v *graph.Graph, u *graph.Graph, p PathMap) {
    var minNode *graph.Node
    var minScore int = MAXWEIGHT

    for edge := range getCut(v, u) {
        score := scoreEdge(edge, p)
        
        if score < minScore {
            minNode = edge.Tail
            minScore = score
        }
    }
    
    if minScore == MAXWEIGHT {
        panic("something bad happened")
    }
    
    v.AddNodes(minNode)
    u.RemoveNodes(minNode)
    p[minNode] = minScore
}

func scoreEdge(edge *graph.Edge, p PathMap) (int) {
    score, _ := p[edge.Head]
    return score + edge.Weight
}

func getCut(src *graph.Graph, dst *graph.Graph) (chan *graph.Edge) {
    edges := make(chan *graph.Edge)
    
    go func() {
        defer close(edges)
        for edge := range src.GetEdges() {
            if dst.Contains(edge.Tail) {
                edges <- edge
            }
        }
    }()
    
    return edges
}