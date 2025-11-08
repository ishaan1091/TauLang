package ast

import (
	"strings"
	"taulang/token"
)

type ConditionalExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (c *ConditionalExpression) TokenLiteral() string {
	return c.Token.Literal
}

func (c *ConditionalExpression) String() string {
	var out strings.Builder

	out.WriteString("if (")
	out.WriteString(c.Condition.String())
	out.WriteString(") ")
	out.WriteString(c.Consequence.String())

	if c.Alternative != nil {
		out.WriteString(" else ")
		out.WriteString(c.Alternative.String())
	}

	return out.String()
}

func (c *ConditionalExpression) expressionNode() {}
