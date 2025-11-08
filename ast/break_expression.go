package ast

import "taulang/token"

type BreakExpression struct {
	Token token.Token
}

func (b *BreakExpression) TokenLiteral() string {
	return b.Token.Literal
}

func (b *BreakExpression) String() string {
	return b.TokenLiteral()
}

func (b *BreakExpression) expressionNode() {}
