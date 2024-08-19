package dot_go

import "fmt"

type EdgeOp string

const (
	DirectedEdgeOp   EdgeOp = "->"
	UndirectedEdgeOp EdgeOp = "--"
)

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
	EdgeOp
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

type NodeItem struct {
	*NodeID
	*SubGraph
}

type Edge struct {
	Lhs NodeItem
	Op  EdgeOp
	Rhs NodeItem
}
