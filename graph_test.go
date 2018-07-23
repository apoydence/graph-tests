package graph_test

import (
	"testing"

	graph "github.com/apoydence/graph-tests"
)

func TestGraph(t *testing.T) {
	t.Parallel()

	g := graph.New()

	va, err := g.AddVertex("A")
	if err != nil {
		t.Fatal(err)
	}

	vb, err := g.AddVertex("B")
	if err != nil {
		t.Fatal(err)
	}

	if len(va.Adjacency()) != 0 {
		t.Fatalf("expected Adjacency list to be empty: %d", len(va.Adjacency()))
	}

	if len(vb.Adjacency()) != 0 {
		t.Fatalf("expected Adjacency list to be empty: %d", len(vb.Adjacency()))
	}

	g.AddDirectionalEdge(va, vb, 0)

	if len(va.Adjacency()) != 1 {
		t.Fatalf("expected Adjacency list to have 1: %d", len(va.Adjacency()))
	}

	if len(vb.Adjacency()) != 0 {
		t.Fatalf("expected Adjacency list to be empty: %d", len(vb.Adjacency()))
	}
}

func TestGraphDuplicateVertices(t *testing.T) {
	t.Parallel()

	g := graph.New()

	_, err := g.AddVertex("A")
	if err != nil {
		t.Fatal(err)
	}

	_, err = g.AddVertex("A")
	if err == nil {
		t.Fatal("expected an error for a duplicate vertex name")
	}
}
