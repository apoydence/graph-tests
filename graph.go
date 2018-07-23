package graph

import "fmt"

type Graph struct {
	vs map[string]*Vertex
}

func New() *Graph {
	return &Graph{
		vs: make(map[string]*Vertex),
	}
}

func (g *Graph) AddVertex(name string) (*Vertex, error) {
	if _, ok := g.vs[name]; ok {
		return nil, fmt.Errorf("duplicate vertex %s", name)
	}

	v := &Vertex{
		neighbors: make(map[*Vertex]int),
	}
	g.vs[name] = v
	return v, nil
}

func (g *Graph) AddDirectionalEdge(va, vb *Vertex, weight int) {
	va.neighbors[va] = weight
}

type Vertex struct {
	neighbors map[*Vertex]int
}

func (v *Vertex) Adjacency() map[*Vertex]int {
	m := make(map[*Vertex]int)
	for k, v := range v.neighbors {
		m[k] = v
	}
	return m
}
