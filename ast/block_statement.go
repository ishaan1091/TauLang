package ast

import (
	"strings"
	"taulang/token"
)

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (b *BlockStatement) TokenLiteral() string {
	return b.Token.Literal
}

func (b *BlockStatement) String() string {
	var out strings.Builder

	out.WriteString("{\n")
	for _, statement := range b.Statements {
		out.WriteString("\t" + statement.String())
	}
	out.WriteString("\n}")

	return out.String()
}

func (b *BlockStatement) statementNode() {}
