package graph_test

import (
	"fmt"

	"github.com/batiazinga/graph"
)

// digraph is a directed graph implementing the Forward interface.
type digraph struct {
	// map from vertices to out edges
	vertices map[string][]string
	// map from edges to target vertices
	edges map[string]string
}

func (g digraph) OutEdges(v string) []string    { return g.vertices[v] }
func (g digraph) NextVertex(v, e string) string { return g.edges[e] }

// bfsVisitorDistance computes the distance between the source and any reachable vertex.
// The distance between two adjacent vertices is one.
type bfsVisitorDistance struct {
	graph.BfsVisitor

	// vertex that is being examined
	// it is initially the empty vertex
	currentVertex string
	// map storing the distance between the source and other vertices
	// it is initially full of -1
	distance map[string]int
}

func (vis *bfsVisitorDistance) DiscoverVertex(v string) {
	if vis.currentVertex != "" {
		vis.distance[v] = vis.distance[vis.currentVertex] + 1

	} else {
		// this is the beginning of the visit
		// v is the source vertex so the distance is 0
		vis.distance[v] = 0
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
		vertices: map[string][]string{
			"A": []string{"AB", "AC"},
			"B": []string{"BD"},
			"C": []string{"CE"},
			"D": []string{"DE"},
			"E": nil,
		},
		edges: map[string]string{
			"AB": "B",
			"AC": "C",
			"BD": "D",
			"CE": "E",
			"DE": "E",
		},
	}

	// create a distance visitor
	// the current vertex is ""
	// and the distance map is initialized to -1 for all vertices
	vis := &bfsVisitorDistance{
		BfsVisitor: graph.BfsVisitorNoOp(),
		distance:   make(map[string]int, len(g.vertices)),
	}
	for v := range g.vertices {
		vis.distance[v] = -1
	}

	// init the color map
	// it is empty and has enough space allocated for all vertices
	colors := make(graph.ColorMap, len(g.vertices))

	// Run the breadth-first visit
	graph.BreadthFirstVisit(g, "A", vis, colors)

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
