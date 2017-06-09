package graph

// Forward is the interface allowing to navigate forward in a graph.
// The graph can directed or undirected.
type Forward interface {
	// OutEdges return the list of edges leaving the given vertex.
	OutEdges(v string) []string

	// NextVertex returns the vertex that can reached from vertex v and navigating along edge e.
	// If e is not an edge leaving v, the empty vertex is returned.
	NextVertex(v, e string) string
}
