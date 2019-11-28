package graph

// color is a 3-element enum
// used to mark vertices during a visit.
type color uint8

// Colors used to mark vertices during a visit:
//
// - white means not discovered yet,
// - gray means discovered but not examined yet and
// - black means examined
const (
	white color = iota // default
	gray
	black
)
