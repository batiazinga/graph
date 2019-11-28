package graph

import "container/heap"

// priorityQueue sorts vertex ids by distance, shorter distances first.
// Distances can be updated.
// It uses a distance map.
type priorityQueue struct {
	v     []string       // list of vertices
	index map[string]int // needed to update the priority queue
	dmap  distanceMap
}

func newPriorityQueue(dmap distanceMap) *priorityQueue {
	q := &priorityQueue{
		index: make(map[string]int),
		dmap:  dmap,
	}
	heap.Init(q)
	return q
}

// push adds v to the queue with distance d.
// The distance map is updated.
//
// Do not push twice the same vertex.
func (q *priorityQueue) push(v string, d float64) {
	q.dmap[v] = d
	heap.Push(q, v)
}

// pop extracts the closest vertex from the queue.
func (q *priorityQueue) pop() string {
	return heap.Pop(q).(string) // we know it's a string
}

// update updates the distance of vertex v.
// The distance map is updated.
//
// Use only vertices which are already in the priority queue.
func (q *priorityQueue) update(v string, d float64) {
	q.dmap[v] = d
	heap.Fix(q, q.index[v])
}

func (q *priorityQueue) Len() int { return len(q.v) }

func (q *priorityQueue) Less(i, j int) bool {
	// lower distances have higher priorities
	return q.dmap.distance(q.v[i]) < q.dmap.distance(q.v[j])
}

func (q *priorityQueue) Swap(i, j int) {
	q.index[q.v[i]], q.index[q.v[j]] = q.index[q.v[j]], q.index[q.v[i]]
	q.v[i], q.v[j] = q.v[j], q.v[i]
}

func (q *priorityQueue) Push(x interface{}) {
	vertex := x.(string) // we know it's a string
	q.index[vertex] = len(q.v)
	q.v = append(q.v, vertex)
}

func (q *priorityQueue) Pop() interface{} {
	n := len(q.v)
	vertex := q.v[n-1]
	q.v = q.v[:n-1]
	delete(q.index, vertex)
	return vertex
}
