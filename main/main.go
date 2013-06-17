package main

import (
    "fmt"
    "flag"
    "github.com/mattomatic/dijkstra/graph"
    "github.com/mattomatic/dijkstra/dijkstra"
)

func main() {
    flag.Parse()
    
    if flag.NArg() != 1 {
        panic("Usage: main <filename>")
    }   
    
    filename := flag.Arg(0)
    g := graph.LoadGraph(filename)
    
    paths := dijkstra.Dijkstra(g, g.Search(1))
    
    for node, path := range paths {
        fmt.Println(node.Id, "=", path)
    }
}