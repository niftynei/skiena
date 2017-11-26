package main

import (
	"github.com/niftynei/algos/five"
	//"fmt"
	//"math/rand"
)

func main() {
	graph := five.MakeGraph(five.StartEdges(), 6, false)
	five.DFS(graph, 1)
}

