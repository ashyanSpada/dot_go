package dot_go

import "fmt"

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
