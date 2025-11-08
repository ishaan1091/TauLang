package ast

import (
	"strings"
	"taulang/token"
)

type WhileLoopExpression struct {
	Token     token.Token
	Condition Expression
	Body      *BlockStatement
}

func (w *WhileLoopExpression) TokenLiteral() string {
	return w.Token.Literal
}

func (w *WhileLoopExpression) String() string {
	var out strings.Builder

	out.WriteString("while (")
	out.WriteString(w.Condition.String())
	out.WriteString(") ")
	out.WriteString(w.Body.String())

	return out.String()
}

func (w *WhileLoopExpression) expressionNode() {}
