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
	pred   map[string]string
	target string
}

func (vis *dijkstraVisitorPath) EdgeRelaxed(from, to string) {
	vis.pred[to] = from
}

func (vis *dijkstraVisitorPath) init(source, target string) {
	vis.target = target
	vis.pred[source] = source // source is its own predecessor
}

func (vis *dijkstraVisitorPath) path() (p []string) {
	// init search
	v := vis.target
	pred, found := vis.pred[v]
	if !found {
		return
	}
	p = append(p, v)

	// loop as long as there is a predecessor
	for v != pred {
		v, pred = pred, vis.pred[pred]
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

	// // init visitor and run the dikstra visit
	vis.init("A", "E")
	distance := graph.DijkstraTo(g, vis, "A", "E")

	// read results
	fmt.Printf("Distance from A to E is 0.4: %v\n", distance > 0.39 && distance < 0.41)
	fmt.Printf("Path from A to E is %v", vis.path())

	// Output:
	// Distance from A to E is 0.4: true
	// Path from A to E is [A B D E]
}
