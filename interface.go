package graph

// Forward is the interface allowing to navigate forward in a graph.
// The graph can be directed or undirected.
type Forward interface {
	// NextVertices returns the list of vertices reachable when leaving the vertex v.
	NextVertices(v string) []string
}

// VertexListForward is a Forward graph
// whose vertices can be listed.
type VertexListForward interface {
	Forward

	// Vertices returns the list of vertices of the graph.
	Vertices() []string
}

// WeightForward is a Forward graph
// with float64 weights on its edges.
type WeightForward interface {
	Forward

	// Weight return the weight of the edge.
	Weight(from, to string) float64
}
