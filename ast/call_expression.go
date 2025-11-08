package ast

import (
	"strings"
	"taulang/token"
)

type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (c *CallExpression) TokenLiteral() string {
	return c.Token.Literal
}

func (c *CallExpression) String() string {
	var out strings.Builder

	out.WriteString(c.Function.String())
	out.WriteString("(")
	for idx, arg := range c.Arguments {
		if idx != 0 {
			out.WriteString(", ")
		}
		out.WriteString(arg.String())
	}
	out.WriteString(")")

	return out.String()
}

func (c *CallExpression) expressionNode() {}
