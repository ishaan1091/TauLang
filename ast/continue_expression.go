package ast

import "taulang/token"

type ContinueExpression struct {
	Token token.Token
}

func (c *ContinueExpression) TokenLiteral() string {
	return c.Token.Literal
}

func (c *ContinueExpression) String() string {
	return c.TokenLiteral()
}

func (c *ContinueExpression) expressionNode() {}
