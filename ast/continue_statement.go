package ast

import "taulang/token"

type ContinueStatement struct {
	Token token.Token
}

func (c *ContinueStatement) TokenLiteral() string {
	return c.Token.Literal
}

func (c *ContinueStatement) String() string {
	return c.TokenLiteral() + ";"
}

func (c *ContinueStatement) statementNode() {}
