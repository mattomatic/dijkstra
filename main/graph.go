package main

type Graph struct {
    nodes []*Node
}

type Edge struct {
	Head *Node
	Tail *Node
	Weight int
}

type Node struct {
	Id int
	Visited bool
	edges []*Edge
}

func NewGraph() *Graph {
    return &Graph{
        make([]*Node, 0),
    }
}

func NewNode(id int) *Node {
    return &Node {
        Id: id,
        edges: make([]*Edge, 0),
    }
}

func NewEdge(head, tail *Node, weight int) *Edge {
    return &Edge {
        head,
        tail,
        weight,
    }
}

func (n *Node) AddEdges(edges ...*Edge) {
    for _, edge := range edges {
        n.edges = append(n.edges, edge)
    }
}

func (n *Node) GetEdges() (chan *Edge) {
    edges := make(chan *Edge)
    
    go func() {
        defer close(edges)
        for _, edge := range n.edges {
            edges <- edge
        }
    }()
    
    return edges
}

func (g *Graph) AddNodes(nodes ...*Node) {
    for _, node := range nodes {
        g.nodes = append(g.nodes, node)
    }
}

func (g *Graph) GetNodes() (chan *Node) {
    nodes := make(chan *Node)
    
    go func() {
        defer close(nodes)
        for _, node := range g.nodes {
            nodes <- node
        }
    }()
    
    return nodes
}

func (g *Graph) Search(id int) *Node {
    for _, node := range g.nodes {
        if node.Id == id {
            return node
        }
    }
    
    return nil
}