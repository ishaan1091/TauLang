package ast

import (
	"strings"
	"taulang/token"
)

type HashLiteral struct {
	Token token.Token
	Pairs []HashPair
}

type HashPair struct {
	Key   Expression
	Value Expression
}

func (h *HashLiteral) TokenLiteral() string {
	return h.Token.Literal
}

func (h *HashLiteral) String() string {
	var out strings.Builder

	var pairs []string
	for _, pair := range h.Pairs {
		pairs = append(pairs, pair.Key.String()+":"+pair.Value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

func (h *HashLiteral) expressionNode() {}
