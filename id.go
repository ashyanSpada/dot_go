package dot_go

import "fmt"

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

func NewNodeID(id ID, port *Port) NodeID {
	return NodeID{
		ID:   id,
		Port: port,
	}
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

func NewPort(id *ID, pt *CompassPt) *Port {
	return &Port{
		ID:        id,
		CompassPt: pt,
	}
}

type ID string

func NewID(a string) ID {
	return ID(a)
}

// ID '=' ID
type IDPair [2]ID

func (i IDPair) IsStmt() {}

func (i IDPair) String() string {
	return fmt.Sprintf("%v = %v", i[0], i[1])
}

func NewIDPair(a, b string) IDPair {
	return [2]ID{
		NewID(a),
		NewID(b),
	}
}
