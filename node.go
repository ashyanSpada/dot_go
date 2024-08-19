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

type Node = NodeStmt

type NodeOption func(*NodeStmt)

func WithAttrs(attrList *AttrList) NodeOption {
	return func(ns *NodeStmt) {
		ns.AttrList = attrList
	}
}

func NewNode(nodeID NodeID, options ...NodeOption) Node {
	s := NodeStmt{
		NodeID: nodeID,
	}
	for _, option := range options {
		option(&s)
	}
	return s
}
