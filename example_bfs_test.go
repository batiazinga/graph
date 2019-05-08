package graph_test

import (
	"fmt"

	"github.com/batiazinga/graph"
	"github.com/batiazinga/graph/visitor"
)

// digraph is a directed graph implementing the Forward interface.
type digraph map[string][]string

func (g digraph) NextVertices(v string) []string { return g[v] }

// bfsVisitorDistance computes the distance between the source and any reachable vertex.
// The distance between two adjacent vertices is one.
//
// Beware that distance from source vertex to non-reachable vertices is zero with this visitor.
type bfsVisitorDistance struct {
	visitor.BfsNoOp // bfsVisitorDistance implement BfsVisitor

	// map storing the distance between the source and other vertices
	// it should initially be empty
	distance map[string]int
}

func (vis *bfsVisitorDistance) TreeEdge(from, to string) {
	vis.distance[to] = vis.distance[from] + 1
}

func ExampleBreadthFirstVisit() {
	// create the following digraph
	// A -> B -> D
	//   \-> C -- \-> E
	g := digraph{
		"A": []string{"B", "C"},
		"B": []string{"D"},
		"C": []string{"E"},
		"D": []string{"E"},
	}

	// create a distance visitor
	vis := &bfsVisitorDistance{
		BfsNoOp:  visitor.BfsNoOp{},
		distance: make(map[string]int),
	}

	// Run the breadth-first visit
	graph.BreadthFirstVisit(g, vis, "A")

	// read results
	fmt.Println("A", vis.distance["A"])
	fmt.Println("B", vis.distance["B"])
	fmt.Println("C", vis.distance["C"])
	fmt.Println("D", vis.distance["D"])
	fmt.Println("E", vis.distance["E"])

	// Output:
	// A 0
	// B 1
	// C 1
	// D 2
	// E 2
}
