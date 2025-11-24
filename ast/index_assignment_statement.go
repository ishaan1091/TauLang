package ast

import (
	"strings"
	"taulang/token"
)

type IndexAssignmentStatement struct {
	Token             token.Token
	IndexedExpression Expression
	Index             Expression
	Value             Expression
}

func (a *IndexAssignmentStatement) TokenLiteral() string {
	return a.Token.Literal
}

func (a *IndexAssignmentStatement) String() string {
	var out strings.Builder

	out.WriteString(a.IndexedExpression.String())
	out.WriteString("[")
	out.WriteString(a.Index.String())
	out.WriteString("]")
	out.WriteString(" = ")
	out.WriteString(a.Value.String())
	out.WriteString(";")

	return out.String()
}

func (a *IndexAssignmentStatement) statementNode() {}
