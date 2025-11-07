package ast

import "taulang/token"

type String struct {
	Token token.Token
	Value string
}

func (s *String) TokenLiteral() string {
	return s.Token.Literal
}

func (s *String) String() string {
	return s.Value
}

func (s *String) expressionNode() {}
