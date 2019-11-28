// Package visitor provides visitors to use with graph algorithms.
package visitor

// BfsNoOp is a BfsVisitor which does nothing.
type BfsNoOp struct{}

func (v BfsNoOp) DiscoverVertex(string)      {}
func (v BfsNoOp) ExamineVertex(string)       {}
func (v BfsNoOp) ExamineEdge(string, string) {}
func (v BfsNoOp) TreeEdge(string, string)    {}
func (v BfsNoOp) NonTreeEdge(string, string) {}
func (v BfsNoOp) GrayTarget(string, string)  {}
func (v BfsNoOp) BlackTarget(string, string) {}
func (v BfsNoOp) FinishVertex(string)        {}

// DfsNoOp is a DfsVisitor which does nothing.
type DfsNoOp struct{}

func (v DfsNoOp) InitializeVertex(string)         {}
func (v DfsNoOp) DiscoverVertex(string)           {}
func (v DfsNoOp) ExamineEdge(string, string)      {}
func (v DfsNoOp) TreeEdge(string, string)         {}
func (v DfsNoOp) BackEdge(string, string)         {}
func (v DfsNoOp) ForwardCrossEdge(string, string) {}
func (v DfsNoOp) FinishVertex(string)             {}

// DijkstraNoOp is a DijkstraVisitor which does nothing.
type DijkstraNoOp struct{}

func (v DijkstraNoOp) DiscoverVertex(string)           {}
func (v DijkstraNoOp) ExamineVertex(string)            {}
func (v DijkstraNoOp) ExamineEdge(string, string)      {}
func (v DijkstraNoOp) EdgeRelaxed(string, string)      {}
func (v DijkstraNoOp) EdgeNotRelaxed(string, string)   {}
func (v DijkstraNoOp) ForwardCrossEdge(string, string) {}
func (v DijkstraNoOp) FinishVertex(string)             {}
