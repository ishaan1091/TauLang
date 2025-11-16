package ast

import (
	"strings"
	"taulang/token"
)

type AssignmentStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (a *AssignmentStatement) TokenLiteral() string {
	return a.Token.Literal
}

func (a *AssignmentStatement) String() string {
	var out strings.Builder

	out.WriteString(a.Name.String())
	out.WriteString(" = ")
	out.WriteString(a.Value.String())

	return out.String()
}

func (a *AssignmentStatement) statementNode() {}
