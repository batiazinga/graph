package graph_test

import (
	"fmt"

	"github.com/batiazinga/graph"
)

// digraph is a directed graph implementing the Forward interface.
type digraph map[string][]string

func (g digraph) NextVertices(v string) []string { return g[v] }

// bfsVisitorDistance computes the distance between the source and any reachable vertex.
// The distance between two adjacent vertices is one.
type bfsVisitorDistance struct {
	graph.BfsVisitorNoOp // bfsVisitorDistance implement BfsVisitor

	// vertex that is being examined
	// it is initially the empty vertex
	currentVertex string
	// map storing the distance between the source and other vertices
	// it should initially be full of -1
	distance map[string]int
}

func (vis *bfsVisitorDistance) DiscoverVertex(v string) {
	if vis.currentVertex == "" {
		// this is the beginning of the visit
		// v is the source vertex so the distance is 0
		vis.distance[v] = 0

	} else {
		// v is discovered from current vertex
		vis.distance[v] = vis.distance[vis.currentVertex] + 1
	}
}
func (vis *bfsVisitorDistance) ExamineVertex(v string) {
	vis.currentVertex = v
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
		"E": nil,
	}

	// create a distance visitor
	// the current vertex is ""
	// and the distance map is initialized to -1 for all vertices
	vis := &bfsVisitorDistance{
		BfsVisitorNoOp: graph.BfsVisitorNoOp{},
		distance:       make(map[string]int, len(g)),
	}
	for v := range g {
		vis.distance[v] = -1
	}

	// Run the breadth-first visit
	graph.BreadthFirstVisit(g, "A", vis)

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
