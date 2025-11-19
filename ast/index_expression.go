package ast

import (
	"strings"
	"taulang/token"
)

type IndexExpression struct {
	Token             token.Token
	IndexedExpression Expression
	Index             Expression
}

func (i *IndexExpression) TokenLiteral() string {
	return i.Token.Literal
}

func (i *IndexExpression) String() string {
	var out strings.Builder

	out.WriteString("(")
	out.WriteString(i.IndexedExpression.String())
	out.WriteString("[")
	out.WriteString(i.Index.String())
	out.WriteString("]")
	out.WriteString(")")

	return out.String()
}

func (i *IndexExpression) expressionNode() {}
