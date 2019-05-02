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
// At some event points the visitor is called.
func DepthFirstVisitFrom(g Forward, vis DfsVisitor, source string) {
	// init color map
	cmap := make(colorMap)

	depthFirstVisitFrom(g, vis, cmap, source)
}

// depthFirstVisitFrom recursively visits g.
func depthFirstVisitFrom(g Forward, vis DfsVisitor, cmap colorMap, source string) {
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
func DepthFirstVisit(g VertexListForward, vis DfsVisitor) {
	// visit vertices and init them
	for _, v := range g.Vertices() {
		vis.InitializeVertex(v)
	}

	// init color map
	cmap := make(colorMap)
	// visit vertices and start a depth-first-visit from each one of them
	for _, v := range g.Vertices() {
		if cmap[v] == white {
			depthFirstVisitFrom(g, vis, cmap, v)
		}

	}
}

// DfsVisitorNoOp is a DfsVisitor which does nothing.
type DfsVisitorNoOp struct{}

func (v DfsVisitorNoOp) InitializeVertex(string)         {}
func (v DfsVisitorNoOp) DiscoverVertex(string)           {}
func (v DfsVisitorNoOp) ExamineEdge(string, string)      {}
func (v DfsVisitorNoOp) TreeEdge(string, string)         {}
func (v DfsVisitorNoOp) BackEdge(string, string)         {}
func (v DfsVisitorNoOp) ForwardCrossEdge(string, string) {}
func (v DfsVisitorNoOp) FinishVertex(string)             {}
