package ast

import "taulang/token"

type BreakStatement struct {
	Token token.Token
}

func (b *BreakStatement) TokenLiteral() string {
	return b.Token.Literal
}

func (b *BreakStatement) String() string {
	return b.TokenLiteral() + ";"
}

func (b *BreakStatement) statementNode() {}
