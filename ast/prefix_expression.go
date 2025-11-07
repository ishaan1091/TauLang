package ast

import (
	"strings"
	"taulang/token"
)

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Operand  Expression
}

func (p *PrefixExpression) TokenLiteral() string {
	return p.Token.Literal
}

func (p *PrefixExpression) String() string {
	var out strings.Builder

	out.WriteString("(")
	out.WriteString(p.Operator)
	out.WriteString(p.Operand.String())
	out.WriteString(")")

	return out.String()
}

func (p *PrefixExpression) expressionNode() {}
