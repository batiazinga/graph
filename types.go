package graph

// VertexID is a string representation of a vertex.
type VertexID string

// EdgeID is string representation of an edge.
type EdgeID string

// Color is a 3-element enum
// used to mark vertices during a visit.
type Color uint8

// Colors used to mark vertices during a visit:
// White means not discovered yet,
// Gray means discovered but not examined yet and
// Black means examined
const (
	White Color = iota // default
	Gray
	Black
)

// VertexColorMap is a map from vertex to color.
// The default color is white so missing is equivalent to white.
type VertexColorMap map[VertexID]Color
