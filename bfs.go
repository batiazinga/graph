package graph

import (
	"container/list"
)

// BfsVisitor is the visitor to be passed to BreadthFirstVisit graph traversal function.
type BfsVisitor interface {
	// DiscoverVertex is called when a new vertex is found.
	DiscoverVertex(v string)

	// ExamineVertex is called when a vertex is dequeued.
	ExamineVertex(v string)

	// ExamineEdge is called when navigating through the edge.
	ExamineEdge(from, to string)

	// TreeEdge is called when navigating to a new vertex.
	// This edge is then an edge of the minimum spanning tree
	// (where the distance between two neighbour vertices is one).
	TreeEdge(from, to string)

	// NonTreeEdge is called when navigating to an already discovered vertex.
	NonTreeEdge(from, to string)

	// GrayTarget is called when the vertex we are navigating to has already been discovered
	// but has not been examined yet.
	GrayTarget(from, to string)

	// BlackTarget is called when the vertex we are navigating to has already been examined.
	BlackTarget(from, to string)

	// FinishVertex is called when a vertex has been examined.
	FinishVertex(v string)
}

// BreadthFirstVisit visits a graph starting from the source vertex
// and visiting closer vertices first.
//
// At some event points the visitor is called.
// An appropriate visitor can then compute distances and shortest paths.
func BreadthFirstVisit(g Forward, vis BfsVisitor, source string) {
	// queue implemented with a list
	queue := list.New()
	// init color map
	cmap := make(colorMap)

	// discover the source vertex:
	// it was white, it is now gray
	vis.DiscoverVertex(source)
	cmap[source] = gray    // mark as discovered
	queue.PushBack(source) // enqueue

	// visit
	for queue.Len() != 0 {
		// dequeue the front element
		// and examine it
		elt := queue.Front()
		queue.Remove(elt)
		v := elt.Value.(string)
		vis.ExamineVertex(v)

		// visit neighbours
		for _, next := range g.NextVertices(v) {
			// leave vertex v toward vertex next
			// and examine the edge
			vis.ExamineEdge(v, next)

			// has next vertex already been discovered
			if cmap[next] == white {
				// vertex has not been discovered yet
				vis.TreeEdge(v, next)
				vis.DiscoverVertex(next)
				cmap[next] = gray    // mark as discovered
				queue.PushBack(next) //enqueue
			} else {
				vis.NonTreeEdge(v, next)
				if cmap[next] == gray {
					vis.GrayTarget(v, next)
				} else {
					vis.BlackTarget(v, next)
				}
			}
		}

		// All neighbours have been found
		// There is nothing left to do with this vertex: turn it to black
		vis.FinishVertex(v)
		cmap[v] = black
	}
}

// BreadthFirstSearch visits a graph starting from the source vertex.
// It visits closer vertices first and stops when the target is discovered.
//
// At some event points the visitor is called.
// An appropriate visitor can then compute distances and shortest paths.
//
// Methods TreeEdge(v, target) and DiscoverVertex(target) are called before the search stops.
// If the target is not reachable from the source, it is equivalent to BreadthFisrtVisit.
func BreadthFirstSearch(g Forward, vis BfsVisitor, source, target string) {
	// queue implemented with a list
	queue := list.New()
	// init color map
	cmap := make(colorMap)

	// discover the source vertex:
	// it was white, it is now gray
	vis.DiscoverVertex(source)
	cmap[source] = gray    // mark as discovered
	queue.PushBack(source) // enqueue

	// visit
	for queue.Len() != 0 {
		// dequeue the front element
		// and examine it
		elt := queue.Front()
		queue.Remove(elt)
		v := elt.Value.(string)
		vis.ExamineVertex(v)

		// visit neighbours
		for _, next := range g.NextVertices(v) {
			// leave vertex v toward vertex next
			// and examine the edge
			vis.ExamineEdge(v, next)

			// has next vertex already been discovered
			if cmap[next] == white {
				// vertex has not been discovered yet
				vis.TreeEdge(v, next)
				vis.DiscoverVertex(next)
				cmap[next] = gray    // mark as discovered
				queue.PushBack(next) //enqueue

				// stop if the new vertex is the target vertex
				if next == target {
					return
				}
			} else {
				vis.NonTreeEdge(v, next)
				if cmap[next] == gray {
					vis.GrayTarget(v, next)
				} else {
					vis.BlackTarget(v, next)
				}
			}
		}

		// All neighbours have been found
		// There is nothing left to do with this vertex: turn it to black
		vis.FinishVertex(v)
		cmap[v] = black
	}
}

// BfsVisitorNoOp is a BfsVisitor which does nothing.
type BfsVisitorNoOp struct{}

func (v BfsVisitorNoOp) DiscoverVertex(string)      {}
func (v BfsVisitorNoOp) ExamineVertex(string)       {}
func (v BfsVisitorNoOp) ExamineEdge(string, string) {}
func (v BfsVisitorNoOp) TreeEdge(string, string)    {}
func (v BfsVisitorNoOp) NonTreeEdge(string, string) {}
func (v BfsVisitorNoOp) GrayTarget(string, string)  {}
func (v BfsVisitorNoOp) BlackTarget(string, string) {}
func (v BfsVisitorNoOp) FinishVertex(string)        {}
