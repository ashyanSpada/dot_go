package dot_go

import (
	"fmt"
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
