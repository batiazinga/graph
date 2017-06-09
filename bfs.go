package graph

import (
	"container/list"
)

// BfsVisitor is the visitor to be passed to BreadthFirstVisit graph traversal function.
type BfsVisitor interface {
	// DiscoverVertex is called when a new (white) vertex is found.
	DiscoverVertex(v string)

	// ExamineVertex is called when a vertex is dequeued.
	ExamineVertex(v string)

	// ExamineEdge is called when navigating through the edge.
	ExamineEdge(e string)

	// TreeEdge is called when navigating to a new (white) vertex.
	// This edge is then an edge of the minimum spanning tree
	// (where the distance between two neighbour vertices is one).
	TreeEdge(e string)

	// NonTreeEdge is called when navigating to an already discovered (gray or black) vertex.
	NonTreeEdge(e string)

	// GrayTarget is called when the vertex we are navigating to has already been discovered
	// but has been examined yet.
	GrayTarget(v string)

	// BlackTarget is called when the vertex we are navigating to has already been examined.
	BlackTarget(v string)

	// FinishVertex is called when a vertex has been examined.
	FinishVertex(v string)
}

// BreadthFirstVisit visits a graph starting from the given source vertex
// and visiting vertices that are closer first.
//
// At some event points the visitor is called.
// An appropriate visitor can then compute shortest paths or precedence map.
//
// A vertex color map must be provided.
// If the whole graph should be visited, the color map should contain only white vertices.
// For performance reason, it should then be empty and have enough space allocated to contain all vertices.
// To avoid visiting some parts of the graph, some vertices may be set to black.
func BreadthFirstVisit(g Forward, source string, vis BfsVisitor, colors ColorMap) {
	// queue implemented with a list
	queue := list.New()

	// discover the source vertex:
	// it was white, it is now gray
	vis.DiscoverVertex(source)
	colors[source] = Gray
	queue.PushBack(source) // enqueue

	// visit
	for queue.Len() != 0 {
		// dequeue the front element
		// and examine it
		elv := queue.Front()
		queue.Remove(elv)
		v := elv.Value.(string)
		vis.ExamineVertex(v)

		// visit neighbours
		for _, e := range g.OutEdges(v) {
			// leave vertex along edge e
			// and examine it
			vis.ExamineEdge(e)

			// find next vertex
			// and test if it has already been discovered
			nextv := g.NextVertex(v, e)
			if colors[nextv] == White {
				// vertex has not been discovered yet
				vis.TreeEdge(e)
				vis.DiscoverVertex(nextv)
				colors[nextv] = Gray
				queue.PushBack(nextv)
			} else {
				vis.NonTreeEdge(e)
				if colors[nextv] == Gray {
					vis.GrayTarget(v)
				} else {
					vis.BlackTarget(v)
				}
			}
		}

		// All neighbours have been found
		// There is nothing left to do with this vertex: turn it to black
		vis.FinishVertex(v)
		colors[v] = Black
	}
}

// TODO: Implement a Breadth-First Search which stops once the target has been found.

// BfsVisitorNoOp returns a BfsVisitor which does nothing.
// It can embedded in a custom BfsVisitor to avoid implementing all empty methods.
func BfsVisitorNoOp() BfsVisitor {
	return bfsVisitorNoOp{}
}

// bfsVisitorNoOp is a BfsVisitor which does nothing.
type bfsVisitorNoOp struct{}

func (v bfsVisitorNoOp) DiscoverVertex(string) {}
func (v bfsVisitorNoOp) ExamineVertex(string)  {}
func (v bfsVisitorNoOp) ExamineEdge(string)    {}
func (v bfsVisitorNoOp) TreeEdge(string)       {}
func (v bfsVisitorNoOp) NonTreeEdge(string)    {}
func (v bfsVisitorNoOp) GrayTarget(string)     {}
func (v bfsVisitorNoOp) BlackTarget(string)    {}
func (v bfsVisitorNoOp) FinishVertex(string)   {}
