package graph

import (
    "bufio"
    "fmt"
    "os"
)

func LoadGraph(filename string) {
    fp, err := os.Open(filename)
    
    if err != nil {
        panic("could not load file!")
    }