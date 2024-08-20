package dot_go

import "fmt"

type EdgeOp string

const (
	DirectedEdgeOp   EdgeOp = "->"
	UndirectedEdgeOp EdgeOp = "--"
)

// edge_lhs : node_id | subgraph
type EdgeLhs struct {
	*NodeID
	*SubGraph
}

func (e EdgeLhs) String() string {
	return fmt.Sprintf(
		"%s %s",
		e.NodeID,
		e.SubGraph,
	)
}

// edge_stmt : edge_lhs edgeRHS [ attr_list ]
type EdgeStmt struct {
	EdgeLhs
	EdgeRhs
	*AttrList
}

func (e EdgeStmt) IsStmt() {}

func (e EdgeStmt) String() string {
	return fmt.Sprintf(
		"%v %v %v",
		e.EdgeLhs,
		e.EdgeRhs,
		SprintPtr(e.AttrList),
	)

}

func (e EdgeStmt) Link(op EdgeOp, next EdgeLhs) EdgeStmt {
	e.EdgeRhs = EdgeRhs{
		EdgeOp:  op,
		EdgeLhs: next,
	}
	return e
}

// edgeRHS	:	edgeop edge_lhs [ edgeRHS ]
type EdgeRhs struct {
	EdgeOp
	EdgeLhs
	*EdgeRhs
}

func (e EdgeRhs) String() string {
	return fmt.Sprintf(
		"%v %v %v",
		e.EdgeOp,
		e.EdgeLhs,
		SprintPtr(e.EdgeRhs),
	)
}

func NewEdgeLhs(nodeID *NodeID, subGraph *SubGraph) EdgeLhs {
	return EdgeLhs{
		NodeID:   nodeID,
		SubGraph: subGraph,
	}
}

func NewEdge(lhs EdgeLhs, op EdgeOp, rhs EdgeLhs) EdgeStmt {
	return EdgeStmt{
		EdgeLhs: lhs,
		EdgeRhs: EdgeRhs{
			EdgeOp:  op,
			EdgeLhs: rhs,
		},
	}
}
