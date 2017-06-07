package graph_test

import (
	"fmt"

	"github.com/batiazinga/graph"
)

// digraph is a directed graph implementing the Forward interface.
type digraph struct {
	vertices map[graph.VertexID][]graph.EdgeID
	edges    map[graph.EdgeID]graph.VertexID
}

func (g digraph) OutEdges(v graph.VertexID) []graph.EdgeID                   { return g.vertices[v] }
func (g digraph) NextVertex(v graph.VertexID, e graph.EdgeID) graph.VertexID { return g.edges[e] }

// bfsVisitorDistance computes the distance between the source and any reachable vertex.
// The distance between two adjacent vertices is one.
type bfsVisitorDistance struct {
	graph.BfsVisitor

	// vertex that is being examined
	// it is initially the empty vertex
	currentVertex graph.VertexID
	// map storing the distance between the source and other vertices
	// it is initially full of -1
	distance map[graph.VertexID]int
}

func (vis *bfsVisitorDistance) DiscoverVertex(v graph.VertexID) {
	if string(vis.currentVertex) != "" {
		vis.distance[v] = vis.distance[vis.currentVertex] + 1

	} else {
		// this is the beginning of the visit
		// v is the source vertex so the distance is 0
		vis.distance[v] = 0
	}
}
func (vis *bfsVisitorDistance) ExamineVertex(v graph.VertexID) {
	vis.currentVertex = v
}

func ExampleBreadthFirstVisit() {
	// create the following digraph
	// A -> B -> D
	//   \-> C -- \-> E
	g := digraph{
		vertices: map[graph.VertexID][]graph.EdgeID{
			graph.VertexID("A"): []graph.EdgeID{graph.EdgeID("AB"), graph.EdgeID("AC")},
			graph.VertexID("B"): []graph.EdgeID{graph.EdgeID("BD")},
			graph.VertexID("C"): []graph.EdgeID{graph.EdgeID("CE")},
			graph.VertexID("D"): []graph.EdgeID{graph.EdgeID("DE")},
			graph.VertexID("E"): nil,
		},
		edges: map[graph.EdgeID]graph.VertexID{
			graph.EdgeID("AB"): graph.VertexID("B"),
			graph.EdgeID("AC"): graph.VertexID("C"),
			graph.EdgeID("BD"): graph.VertexID("D"),
			graph.EdgeID("CE"): graph.VertexID("E"),
			graph.EdgeID("DE"): graph.VertexID("E"),
		},
	}

	// create a distance visitor
	// the current vertex is ""
	// and the distance map is initialized to -1 for all vertices
	vis := &bfsVisitorDistance{
		BfsVisitor: graph.BfsVisitorNoOp(),
		distance:   make(map[graph.VertexID]int, len(g.vertices)),
	}
	for v := range g.vertices {
		vis.distance[v] = -1
	}

	// init the color map
	// it is empty and has enough space allocated for all vertices
	colors := make(graph.VertexColorMap, len(g.vertices))

	// Run the breadth-first visit
	graph.BreadthFirstVisit(g, graph.VertexID("A"), vis, colors)

	// read results
	fmt.Println("A", vis.distance[graph.VertexID("A")])
	fmt.Println("B", vis.distance[graph.VertexID("B")])
	fmt.Println("C", vis.distance[graph.VertexID("C")])
	fmt.Println("D", vis.distance[graph.VertexID("D")])
	fmt.Println("E", vis.distance[graph.VertexID("E")])

	// Output:
	// A 0
	// B 1
	// C 1
	// D 2
	// E 2
}
