package five

import (
	"fmt"
)

func StartEdges() [][]int {
	return [][]int{
		[]int{5,1},
		[]int{1,2},
		[]int{3,4},
		[]int{4,6},
		[]int{5,4},
		[]int{3,2}}
}

type Edgenode struct {
	Y int // adjacency info
	Weight int // if any
	Next *Edgenode
}

type Graph struct {
	Edges []*Edgenode // adjacency info (no matrix!)
	Degree  []int // outdegree of ea vertex ??
	NVertices int
	NEdges int
	Directed bool
}

func (g *Graph) Print() {
	for i := 1; i <= g.NVertices; i++ {
		fmt.Printf("Node %d:", i)
		p := g.Edges[i]
		for p != nil {
			fmt.Printf(" => %d", p.Y)
			p = p.Next
		}
		fmt.Printf("\n")
	}
}

func (g *Graph) insertEdge(input []int, directed bool) {
	node := &Edgenode{}
	x, y := input[0], input[1]

	node.Y = y
	node.Weight = 0
	node.Next = g.Edges[x]
	g.Edges[x] = node
	g.Degree[x] += 1

	if !directed {
		g.insertEdge([]int{y,x}, true)
	} else {
		g.NEdges++
	}
}

func MakeGraph(input [][]int, verticeCount int, directed bool) (*Graph) {
	graph := &Graph{}
	graph.NVertices = verticeCount
	graph.Directed = directed
	graph.Edges = make([]*Edgenode, verticeCount + 1)
	graph.Degree = make([]int, verticeCount + 1)
	for _, edge := range input {
		graph.insertEdge(edge, directed)
	}
	return graph
}

func BFS(g *Graph, start int) []int {
	processed := make([]bool, g.NVertices + 1)
	discovered := make([]bool, g.NVertices + 1)
	parent := make([]int, g.NVertices + 1)

	q := &Queue{}
	q.Enqueue(start)
	discovered[start] = true

	for !q.Empty() {
		v := q.Dequeue()
		processEarly(v)
		processed[v] = true
		p := g.Edges[v]
		for p != nil {
			y := p.Y
			if !processed[y] || g.Directed {
				processEdge(v,y)
			}
			if !discovered[y] {
				q.Enqueue(y)
				discovered[y] = true
				parent[y] = v
			}
			p = p.Next
		}
		processLate(v)

	}

	return parent
}

func processEarly(vertex int) {
	fmt.Printf("Processing vertext %d\n", vertex)
}

func processLate(vertex int) {
}

func processEdge(x,y int) {
	fmt.Printf(" === edge (%d,%d)\n", x,y)
}

func FindPath(start, end int, parents []int) {
	if start == end || end == 0 {
		fmt.Printf("start %d\n", start)
	} else {
		FindPath(start, parents[end], parents)
		fmt.Printf("%d\n", end)
	}
}

func DFS(g *Graph, start int) {
	processed := make([]bool, g.NVertices + 1)
	discovered := make([]bool, g.NVertices + 1)

	s := &Stack{}
	s.Push(start)
	discovered[start] = true

	for !s.Empty() {
		v := s.Pop()
		processEarly(v)
		processed[v] = true
		p := g.Edges[v]
		for p != nil {
			y := p.Y
			if !processed[y] || g.Directed {
				processEdge(v,y)
			}
			if !discovered[y] {
				s.Push(y)
				discovered[y] = true
			}
			p = p.Next
		}
		processLate(v)

	}
}

func DFSr(g *Graph) {
}
