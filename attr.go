package dot_go

import (
	"fmt"
	"strings"
)

// attr_stmt : (graph | node | edge) attr_list
type AttrStmt struct {
	AttrType
	AttrList
}

func (n AttrStmt) IsStmt() {}

func (a AttrStmt) String() string {
	return fmt.Sprintf("%v %v", a.AttrType, a.AttrList)
}

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
