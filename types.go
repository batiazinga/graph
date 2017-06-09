package graph

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

// ColorMap is a map from (vertex) id  to color.
// The default color is white so missing id is equivalent to white.
type ColorMap map[string]Color
