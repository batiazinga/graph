package graph

// Forward is the interface allowing to navigate forward in a graph.
// The graph can be directed or undirected.
type Forward interface {
	// NextVertices returns the list of vertices reachable when leaving the vertex v.
	NextVertices(v string) []string
}
