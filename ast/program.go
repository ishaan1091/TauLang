package ast

import (
	"strings"
)

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out strings.Builder

	for idx, s := range p.Statements {
		out.WriteString(s.String())
		if idx+1 != len(p.Statements) {
			out.WriteString("\n")
		}
	}

	return out.String()
}
