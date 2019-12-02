package graph_test

import (
	"fmt"
	"sort"

	"github.com/batiazinga/graph"
	"github.com/batiazinga/graph/visitor"
)

// distanceGraph is an undirected graph implementing the WeightForward interface.
type distanceGraph struct {
	next     map[string][]string
	distance map[string]float64
}

func (g distanceGraph) NextVertices(v string) []string { return g.next[v] }
func (g distanceGraph) Weight(v, w string) float64 {
	sorted := []string{v, w}
	sort.Strings(sorted)
	return g.distance[sorted[0]+"-"+sorted[1]]
}

// dijkstraVisitorPath computes the path between the source and a target vertex.
type dijkstraVisitorPath struct {
	visitor.DijkstraNoOp // dijkstraVisitorPath implement DijkstraVisitor

	// information to build a path:
	// predecessor map and target
	pred map[string]string
}

func (vis *dijkstraVisitorPath) EdgeRelaxed(from, to string) {
	vis.pred[to] = from
}

func (vis *dijkstraVisitorPath) path(source, target string) (p []string) {
	// init search
	v := target
	pred, found := vis.pred[v]
	if !found {
		return // did not reach target
	}
	p = append(p, v)

	// loop as long as v is in the map (source is not in the map)
	for found {
		v = pred
		pred, found = vis.pred[v]
		p = append(p, v)
	}

	// reverse path
	last := len(p) - 1
	for i := 0; i < len(p)/2; i++ {
		p[i], p[last-i] = p[last-i], p[i]
	}

	return
}

func ExampleDijkstraTo() {
	// create the following graph
	// A -0.1- B -0.2- D -0.1-
	//   \---0.6--- C ---0.3-- \-> E
	g := distanceGraph{
		// make it undirected!
		next: map[string][]string{
			"A": []string{"B", "C"},
			"B": []string{"A", "D"},
			"C": []string{"A", "E"},
			"D": []string{"B", "E"},
			"E": []string{"C", "D"},
		},
		distance: map[string]float64{
			"A-B": 0.1,
			"B-D": 0.2,
			"D-E": 0.1,
			"A-C": 0.6,
			"C-E": 0.3,
		},
	}

	// create a path visitor
	vis := &dijkstraVisitorPath{
		DijkstraNoOp: visitor.DijkstraNoOp{},
		pred:         make(map[string]string),
	}

	// run the dikstra visit
	graph.DijkstraTo(g, vis, "A", "E")

	// read results
	path := vis.path("A", "E")
	var distance float64
	for i := 0; i < len(path)-1; i++ {
		distance += g.Weight(path[i], path[i+1])
	}
	fmt.Printf("Distance from A to E is 0.4: %v\n", distance > 0.39 && distance < 0.41)
	fmt.Printf("Path from A to E is %v", path)

	// Output:
	// Distance from A to E is 0.4: true
	// Path from A to E is [A B D E]
}
