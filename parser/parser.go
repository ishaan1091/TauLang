package parser

import (
	"fmt"
	"taulang/ast"
	"taulang/lexer"
	"taulang/token"
)

type Parser interface {
	Parse() (*ast.Program, error)
	Errors() []string
}

type (
	prefixParseFunction func() ast.Expression
	infixParseFunction  func(ast.Expression) ast.Expression
)

type parser struct {
	lexer  lexer.Lexer
	errors []string

	currToken token.Token
	peekToken token.Token

	prefixParseFunctions map[token.Token]prefixParseFunction
	infixParseFunctions  map[token.Token]infixParseFunction
}

func NewParser(l lexer.Lexer) (Parser, error) {
	p := parser{
		lexer:  l,
		errors: []string{},
	}

	p.prefixParseFunctions = make(map[token.Token]prefixParseFunction)
	p.infixParseFunctions = make(map[token.Token]infixParseFunction)

	// advancing tokens two times to populate both next and curr tokens
	err := p.nextToken()
	if err != nil {
		return nil, err
	}
	err = p.nextToken()
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (p *parser) Parse() (*ast.Program, error) {
	program := ast.Program{}

	program.Statements = []ast.Statement{}

	for p.currToken.Type != token.EOF {
		statement, err := p.parseStatement()
		if err != nil {
			return nil, err
		}

		if statement != nil {
			program.Statements = append(program.Statements)
		}

		err = p.nextToken()
		if err != nil {
			return nil, err
		}
	}

	return &program, nil
}

func (p *parser) Errors() []string {
	return p.errors
}

func (p *parser) nextToken() error {
	p.currToken = p.peekToken

	tok, err := p.lexer.NextToken()
	if err != nil {
		return fmt.Errorf("got error while parsing next token: %w", err)
	}
	p.peekToken = tok
	return nil
}

func (p *parser) parseStatement() (ast.Statement, error) {
	return nil, nil
}
