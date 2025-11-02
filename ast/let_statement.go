package ast

import (
	"strings"
	"taulang/token"
)

// LetStatement represents a variable binding, e.g., let x = 5;
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (l *LetStatement) statementNode() {}

func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

func (l *LetStatement) String() string {
	var out strings.Builder

	out.WriteString("let ")
	if l.Name != nil {
		out.WriteString(l.Name.String())
	}
	out.WriteString(" = ")
	if l.Value != nil {
		out.WriteString(l.Value.String())
	}
	out.WriteString(";")

	return out.String()
}
