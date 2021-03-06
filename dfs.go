package graph

// DfsVisitor is the visitor to be passed to DepthFirstVisit graph traversal function.
type DfsVisitor interface {
	// InitializeVertex is called for each vertex before the visit starts.
	// It is called only for DepthFirstVisit.
	InitializeVertex(v string)

	// DiscoverVertex is called when a new vertex if found.
	DiscoverVertex(v string)

	// ExamineEdge is called when navigating through the edge.
	ExamineEdge(from, to string)

	// TreeEdge is called when navigating to a new vertex.
	// This edge is then an edge of the search tree.
	TreeEdge(from, to string)

	// BackEdge is called when a visited but unfinished vertex is found.
	// On an undirected graph this is called for each tree edge.
	BackEdge(from, to string)

	// ForwardCrossEdge is called when a finished vertex is found.
	// This is never called on an undirected graph.
	ForwardCrossEdge(from, to string)

	// FinishVertex is called when all out edges have been added to the search tree
	// and all corresponding adjacent vertices are finished.
	FinishVertex(v string)
}

// DepthFirstVisitFrom performs a depth-first-search from the source vertex.
// When possible, it chooses a vertex adjacent to the current vertex to visit next.
// Otherwise it backtracks to the last vertex with unvisited adjacent vertices.
//
// The slice returned by calls to NextVertices is never modified.
// So there is no risk of accidentally modifying g.
//
// At some event points the visitor is called.
func DepthFirstVisitFrom(g Forward, vis DfsVisitor, source string) {
	// init color map
	cmap := make(map[string]color)

	depthFirstVisitFrom(g, vis, cmap, source)
}

// depthFirstVisitFrom recursively visits g.
func depthFirstVisitFrom(g Forward, vis DfsVisitor, cmap map[string]color, source string) {
	// Discover the source vertex and turn it to gray
	vis.DiscoverVertex(source)
	cmap[source] = gray

	// visit out edges and adjacent vertices
	for _, next := range g.NextVertices(source) {
		vis.ExamineEdge(source, next)

		switch cmap[next] {
		case white:
			vis.TreeEdge(source, next)
			depthFirstVisitFrom(g, vis, cmap, next)
		case gray:
			vis.BackEdge(source, next)
		case black:
			vis.ForwardCrossEdge(source, next)
		}
	}

	// all adjacent vertices have been discovered
	// finish this vertex and backtrack
	vis.FinishVertex(source)
	cmap[source] = black
}

// DepthFirstVisit is similar to DepthFirstVisitFrom but it visits the whole graph.
// It needs a graph whose vertices can be listed.
//
// The slices returned by calls to NextVertices and Vertices are never modified.
// So there is no risk of accidentally modifying g.
func DepthFirstVisit(g VertexListForward, vis DfsVisitor) {
	// visit vertices and init them
	for _, v := range g.Vertices() {
		vis.InitializeVertex(v)
	}

	// init color map
	cmap := make(map[string]color)
	// visit vertices and start a depth-first-visit from each one of them
	for _, v := range g.Vertices() {
		if cmap[v] == white {
			depthFirstVisitFrom(g, vis, cmap, v)
		}

	}
}
