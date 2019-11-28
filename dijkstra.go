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
	// of if it just discovered.
	EdgeRelaxed(from, to string)

	// EdgeNotRelaxed is called when a longer path to vertex 'to' is found.
	EdgeNotRelaxed(from, to string)

	// FinishVertex is called when a vertex has been examined.
	FinishVertex(v string)
}

// Dijkstra returns the distance map from the source vertex to any vertex in the graph.
//
// It works for both undirected and directed graphs with non-negative distances and is non destructive.
// If some vertices cannot be reached from the source they are not in the distance map.
//
// The slice returned by calls to NextVertices is never modified.
// So there is no risk of accidentally modifying g.
//
// If all weights are equal to one, use breadth-first-search with the appropriate visitor instead.
func Dijkstra(g WeightForward, vis DijkstraVisitor, source string) map[string]float64 {
	dmap := dijkstra(g, vis, source, nil)
	return map[string]float64(dmap)
}

// DijkstraTo returns the distance from the source vertex to the target vertex.
//
// It works for both undirected and directed graphs with non-negative distances and is non destructive.
// It returns +Inf if the target cannot be reached from the source.
//
// A visitor can be used to collect more information, for example the path from the source to the target
// or the number of visited vertices.
//
// If all weights are equal to one, use breadth-first-search with the appropriate visitor instead.
func DijkstraTo(g WeightForward, vis DijkstraVisitor, source, target string) float64 {
	dmap := dijkstra(g, vis, source, &target)
	return dmap.distance(target)
}

func dijkstra(g WeightForward, vis DijkstraVisitor, source string, target *string) distanceMap {
	// init queue, color map and distance map
	cmap := make(colorMap)
	dist := make(distanceMap)
	queue := newPriorityQueue(dist)

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
		v := queue.pop()
		vis.ExamineVertex(v)

		// stop here if the target vertex has been found
		if lookForTarget && v == *target {
			return dist
		}

		// visit neighbours
		for _, next := range g.NextVertices(v) {
			vis.ExamineEdge(v, next)

			// if already visited, ignore it
			// however we do not need to test this case
			// - if graph is undirected, this never happens
			// - if graph is directed, we have just found a longer path to next

			tentative := dist.distance(v) + g.Weight(v, next)
			if tentative < dist.distance(next) {
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

	return dist
}
