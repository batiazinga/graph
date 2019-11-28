package graph

import "math"

// distanceMap maps vertex IDs to distances (float64).
type distanceMap map[string]float64

// distance returns the distance associated to vertex v.
// It returns +Inf if v is unknown.
func (m distanceMap) distance(v string) float64 {
	d, ok := m[v]
	if !ok {
		return math.Inf(1)
	}
	return d
}
