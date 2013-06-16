package main

import (
    "fmt"
    "flag"
)

func main() {
    flag.Parse()
    filename := flag.Arg(0)
    g := LoadGraph(filename)
    
    for node := range g.GetNodes() {
        fmt.Println(node)
    }
}