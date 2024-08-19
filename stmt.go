package dot_go

import (
	"fmt"
	"strings"
)

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
