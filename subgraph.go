package dot_go

import "fmt"

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
