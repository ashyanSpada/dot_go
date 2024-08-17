package dot_go

import (
	"fmt"
	"strings"
)

type Keyword uint32

const (
	KeywordStrict Keyword = iota
	KeywordGraph
	KeywordDigraph
	KeywordNode
	KeywordEdge
	KeywordSubgraph
)

type EdgeOp string

const (
	DirectedEdgeOp   EdgeOp = "->"
	UndirectedEdgeOp EdgeOp = "--"
)

type GType string

const (
	GraphGType   GType = "graph"
	DigraphGType GType = "digraph"
)

type AttrType string

const (
	GraphAttrType AttrType = "graph"
	NodeAttrType  AttrType = "node"
	EdgeAttrType  AttrType = "edge"
)

// compass_pt :	n | ne | e | se | s | sw | w | nw | c | _)
type CompassPt string

const (
	CompassPtN  CompassPt = "n"
	CompassPtNe CompassPt = "ne"
	CompassPtE  CompassPt = "e"
	CompassPtSE CompassPt = "se"
	CompassPtS  CompassPt = "s"
	CompassPtSW CompassPt = "sw"
	CompassPtW  CompassPt = "w"
	CompassPtNW CompassPt = "nw"
)

// graph : [ strict ] (graph | digraph) [ ID ] '{' stmt_list '}'
type Graph struct {
	strict bool
	GType
	*ID
	StmtList
}

func (g *Graph) AddStmt(stmt Stmt) *Graph {
	g.StmtList = append(g.StmtList, stmt)
	return g
}

func (g Graph) String() string {
	var ans string
	if g.strict {
		ans += "strict "
	}
	ans += fmt.Sprintf(
		"%v %v {\n%v\n}",
		g.GType,
		SprintPtr(g.ID),
		g.StmtList,
	)
	return ans
}

// stmt_list : [ stmt [ ';' ] stmt_list ]
type StmtList []Stmt

func (s StmtList) String() string {
	ans := make([]string, 0, len(s))
	for _, stmt := range s {
		ans = append(ans, fmt.Sprintf("%v", stmt))
	}
	return strings.Join(ans, ";\n")
}

// Stmt : node_stmt |  edge_stmt | attr_stmt | id_pair | subgraph
type Stmt interface {
	// identify a stmt
	IsStmt()
}

// attr_stmt : (graph | node | edge) attr_list
type AttrStmt struct {
	AttrType
	AttrList
}

func (n AttrStmt) IsStmt() {}

func (a AttrStmt) String() string {
	return fmt.Sprintf("%v %v", a.AttrType, a.AttrList)
}

// node_stmt : node_id [ attr_list ]
type NodeStmt struct {
	NodeID
	*AttrList
}

func (n NodeStmt) IsStmt() {}

func (n NodeStmt) String() string {
	return fmt.Sprintf(
		"%v %v",
		n.NodeID,
		SprintPtr(n.AttrList),
	)
}

// edge_stmt : (node_id | subgraph) edgeRHS [ attr_list ]
type EdgeStmt struct {
	*NodeID
	*SubGraph
	EdgeRhs
	*AttrList
}

func (e EdgeStmt) IsStmt() {}

func (e EdgeStmt) String() string {
	return fmt.Sprintf(
		"%v %v %v %v",
		SprintPtr(e.NodeID),
		SprintPtr(e.SubGraph),
		e.EdgeRhs,
		SprintPtr(e.AttrList),
	)

}

// edgeRHS	:	edgeop (node_id | subgraph) [ edgeRHS ]
type EdgeRhs struct {
	EdgeOp string
	*NodeID
	*SubGraph
	*EdgeRhs
}

func (e EdgeRhs) String() string {
	return fmt.Sprintf(
		"%v %v %v %v",
		e.EdgeOp,
		SprintPtr(e.NodeID),
		SprintPtr(e.SubGraph),
		SprintPtr(e.EdgeRhs),
	)
}

// node_id : ID [ port ]
type NodeID struct {
	ID
	*Port
}

func (n NodeID) String() string {
	return fmt.Sprintf(
		"%v %v",
		n.ID,
		SprintPtr(n.Port),
	)
}

// port	: ':' ID [ ':' compass_pt ] | ':' compass_pt
type Port struct {
	*ID
	*CompassPt
}

func (p Port) String() string {
	return fmt.Sprintf(
		": %v %v",
		SprintPtr(p.ID),
		SprintPtr(p.CompassPt),
	)
}

type ID string

// attr_list : '[' [ a_list ] ']' [ attr_list ]
type AttrList []AList

func (a AttrList) String() string {
	ans := make([]string, 0, len(a))
	for _, item := range ans {
		ans = append(ans, fmt.Sprintf("[ %v ]", item))
	}
	return strings.Join(ans, " ")
}

// a_list :	ID '=' ID [ (';' | ',') ] [ a_list ]
type AList []IDPair

func (a AList) String() string {
	ans := make([]string, 0, len(a))
	for _, pair := range ans {
		ans = append(ans, fmt.Sprintf("%v", pair))
	}
	return strings.Join(ans, ";")
}

// subgraph	: [ subgraph [ ID ] ] '{' stmt_list '}'
type SubGraph struct {
	*ID
	StmtList
}

func (s *SubGraph) AddStmt(stmt Stmt) *SubGraph {
	s.StmtList = append(s.StmtList, stmt)
	return s
}

func (s SubGraph) IsStmt() {}

func (s SubGraph) String() string {
	return fmt.Sprintf(
		"subgraph %v {\n%v\n}",
		SprintPtr(s.ID),
		s.StmtList,
	)
}

// ID '=' ID
type IDPair [2]ID

func (i IDPair) IsStmt() {}

func (i IDPair) String() string {
	return fmt.Sprintf("%v = %v", i[0], i[1])
}

func SprintPtr[T any](input *T) string {
	if input == nil {
		return ""
	}
	return fmt.Sprintf("%v", *input)
}
