package parser

import (
	"go/token"
	"taulang/ast"
	"taulang/lexer"
)

type Parser interface {
	Parse() *ast.Program
}

type (
	prefixParseFunction func() ast.Expression
	infixParseFunction  func(ast.Expression) ast.Expression
)

type parser struct {
	lexer                *lexer.Lexer
	errors               []string
	prefixParseFunctions map[token.Token]prefixParseFunction
	infixParseFunctions  map[token.Token]infixParseFunction
}

func NewParser(l *lexer.Lexer) Parser {
	return &parser{
		lexer:  l,
		errors: []string{},
	}
}

func (p *parser) Parse() *ast.Program {
	return nil
}
