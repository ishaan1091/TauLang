package ast

type Program struct {
	Statements []Statement
}

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node

	// Dummy method to distinguish Statement from Expression
	// else without this method, any struct implementing Node would be both
	// Statement and Expression, which is not desired.
	statementNode()
}

type Expression interface {
	Node

	// Dummy method to distinguish Expression from Statement
	// else without this method, any struct implementing Node would be both
	// Statement and Expression, which is not desired.
	expressionNode()
}
