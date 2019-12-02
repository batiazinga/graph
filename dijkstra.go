package graph

// DijkstraVisitor is the visitor to be passed to Dijkstra functions.
type DijkstraVisitor interface {
	// DiscoverVertex is called when a new vertex is found.
	DiscoverVertex(v string)

	// ExamineVertex is called when a vertex is dequeued.
	ExamineVertex(v string)

	// ExamineEdge is called when navigating through the edge.
	ExamineEdge(from, to string)

	// EdgeRelaced is called when a shorter path to vertex 'to' is found
	// or if it was just discovered.
	EdgeRelaxed(from, to string)

	// EdgeNotRelaxed is called when a longer path to vertex 'to' is found.
	EdgeNotRelaxed(from, to string)

	// FinishVertex is called when a vertex has been examined.
	FinishVertex(v string)
}

// Dijkstra visits the graph in Dijkstra order, i.e. closest vertices first.
// It stops when all vertices reachable from the source have been visited.
//
// Shortest paths and distances can be computed thanks to an appropriate visitor.
//
// It works for both undirected and directed graphs with non-negative distances and is non destructive.
//
// The slice returned by calls to NextVertices is never modified.
// So there is no risk of accidentally modifying g.
//
// If all weights are equal to one, use breadth-first-search with the appropriate visitor instead.
func Dijkstra(g WeightForward, vis DijkstraVisitor, source string) {
	dijkstra(g, vis, source, nil)
}

// DijkstraTo visits the graph in Dijkstra order, i.e. closest vertices first.
// It stops when the target vertex has been found or
// when all vertices reachable from the source have been visited.
//
// Shortest path and distance can be computed thanks to an appropriate visitor.
//
// It works for both undirected and directed graphs with non-negative distances and is non destructive.
// It returns +Inf if the target cannot be reached from the source.
//
// A visitor can be used to collect more information, for example the path from the source to the target
// or the number of visited vertices.
//
// If all weights are equal to one, use breadth-first-search with the appropriate visitor instead.
func DijkstraTo(g WeightForward, vis DijkstraVisitor, source, target string) {
	dijkstra(g, vis, source, &target)
}

func dijkstra(g WeightForward, vis DijkstraVisitor, source string, target *string) {
	// init queue, color map and distance map
	cmap := make(map[string]color)
	queue := newPriorityQueue()

	// discover the source vertex:
	// it was white, it is now gray
	vis.DiscoverVertex(source)
	cmap[source] = gray     // mark as discovered
	queue.push(source, 0.0) // enqueue, distance from source to itself is zero

	// confiture visit: are we looking for target?
	lookForTarget := target != nil
	// visit
	for queue.Len() != 0 {
		// pop closest vertex and examine it
		v, d := queue.pop()
		vis.ExamineVertex(v)

		// stop here if the target vertex has been found
		if lookForTarget && v == *target {
			return
		}

		// visit neighbours
		for _, next := range g.NextVertices(v) {
			vis.ExamineEdge(v, next)

			// if already visited, ignore it
			if cmap[next] == black {
				continue
			}

			tentative := d + g.Weight(v, next)
			if tentative < queue.distance(next) {
				// a shorter path to next has been found
				vis.EdgeRelaxed(v, next)
				if cmap[next] == white {
					vis.DiscoverVertex(next)
					queue.push(next, tentative)
					cmap[next] = gray
				} else if cmap[next] == gray {
					queue.update(next, tentative)
				}
			} else {
				// found a longer path to next
				vis.EdgeNotRelaxed(v, next)
			}
		}

		vis.FinishVertex(v)
		cmap[v] = black
	}

	return
}
