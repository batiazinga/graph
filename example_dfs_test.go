package graph_test

import (
	"fmt"
	"sort"

	"github.com/batiazinga/graph"
)

// vertexListDAG is a directed graph implementing the VertexListForward interface.
type vertexListDAG map[string][]string

func (g vertexListDAG) NextVertices(v string) []string {
	// deterministic order to make the visit deterministic
	return g[v]
}

func (g vertexListDAG) Vertices() []string {
	vertices := make([]string, 0, len(g))
	for v := range g {
		vertices = append(vertices, v)
	}
	// sort to make order deterministic
	// the visit can then be deterministic
	sort.Strings(vertices)
	return vertices
}

// toposortVisitor computes a topological ordering of a graph.
// It is a DfsVisitor
type toposortVisitor struct {
	graph.DfsVisitorNoOp // toposortVisitor implement BfsVisitor

	// reverse topological order
	order []string
}

func (vis *toposortVisitor) FinishVertex(v string) {
	vis.order = append(vis.order, v)
}

func ExampleDepthFirstVisit() {
	g := vertexListDAG{
		"A": []string{"B", "C"},
		"B": []string{"D", "E"},
		"C": []string{"E"},
		"D": []string{"E"},
		"E": nil,
		"F": []string{"B"},
	}

	// create a distance visitor
	vis := &toposortVisitor{
		DfsVisitorNoOp: graph.DfsVisitorNoOp{},
	}

	// Run the breadth-first visit
	graph.DepthFirstVisit(g, vis)

	// read results
	for _, v := range vis.order {
		fmt.Println(v)
	}

	// Output:
	// E
	// D
	// B
	// C
	// A
	// F
}
