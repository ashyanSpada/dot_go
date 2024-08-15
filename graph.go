package dot_go

type Keyword uint32

const (
	KeywordStrict Keyword = iota
	KeywordGraph
	KeywordDigraph
	KeywordNode
	KeywordEdge
	KeywordSubgraph
)

type Graph struct {
	stmts  []Stmt
	strict bool
	gType  int
}

func (g *Graph) AddStmt(stmt Stmt) *Graph {
	g.stmts = append(g.stmts, stmt)
	return g
}

type Stmt interface{}

type NodeStmt struct{}

type EdgeStmt struct{}

type AttrStmt struct {
}

type SimpleStmt struct{}

type SubGraph struct {
	stmts []Stmt
}

func (s *SubGraph) AddStmt(stmt Stmt) *SubGraph {
	s.stmts = append(s.stmts, stmt)
	return s
}
