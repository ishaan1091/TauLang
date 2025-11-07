package parser

import (
	"fmt"
	"taulang/ast"
	"taulang/lexer"
	"taulang/token"
)

type Parser interface {
	Parse() *ast.Program
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

	prefixParseFunctions map[token.Type]prefixParseFunction
	infixParseFunctions  map[token.Type]infixParseFunction
}

func NewParser(l lexer.Lexer) Parser {
	p := parser{
		lexer:  l,
		errors: []string{},
	}

	p.prefixParseFunctions = make(map[token.Type]prefixParseFunction)
	p.infixParseFunctions = make(map[token.Type]infixParseFunction)

	// advancing tokens two times to populate both next and curr tokens
	p.nextToken()
	p.nextToken()

	return &p
}

func (p *parser) Parse() *ast.Program {
	program := ast.Program{}

	program.Statements = []ast.Statement{}

	for p.currToken.Type != token.EOF {
		if statement := p.parseStatement(); statement != nil {
			program.Statements = append(program.Statements, statement)
		}

		p.nextToken()
	}

	return &program
}

func (p *parser) Errors() []string {
	return p.errors
}

func (p *parser) nextToken() {
	p.currToken = p.peekToken

	tok := p.lexer.NextToken()
	p.peekToken = tok
}

func (p *parser) currTokenIs(tok token.Type) bool {
	return p.currToken.Type == tok
}

func (p *parser) peekTokenIs(tok token.Type) bool {
	return p.peekToken.Type == tok
}

func (p *parser) expectPeekToken(tok token.Type) bool {
	if p.peekTokenIs(tok) {
		p.nextToken()
		return true
	}

	p.peekTokenMismatchError(tok)
	return false
}

func (p *parser) peekTokenMismatchError(expected token.Type) {
	actual := p.peekToken
	msg := fmt.Sprintf("expected next token to be %s, got %s", expected, actual.Type)

	// Append token literal in case of illegal tokens to give more visibility into the error
	if actual.Type == token.ILLEGAL && actual.Literal != "" {
		msg += fmt.Sprintf(" (%s)", actual.Literal)
	}

	p.errors = append(p.errors, msg)
}

func (p *parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *parser) parseLetStatement() ast.Statement {
	statement := ast.LetStatement{Token: p.currToken}

	if !p.expectPeekToken(token.IDENTIFIER) {
		return nil
	}
	statement.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

	if !p.expectPeekToken(token.ASSIGNMENT) {
		return nil
	}

	// TODO: Parse expression
	for !p.peekTokenIs(token.SEMICOLON) && !p.peekTokenIs(token.EOF) {
		p.nextToken()
	}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return &statement
}

func (p *parser) parseReturnStatement() ast.Statement {
	statement := ast.ReturnStatement{Token: p.currToken}

	// TODO: Parse expression
	for !p.peekTokenIs(token.SEMICOLON) && !p.peekTokenIs(token.EOF) {
		p.nextToken()
	}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return &statement
}

func (p *parser) parseExpressionStatement() ast.Statement {
	statement := ast.ExpressionStatement{Token: p.currToken}

	// TODO: Parse expression
	for !p.peekTokenIs(token.SEMICOLON) && !p.peekTokenIs(token.EOF) {
		p.nextToken()
	}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return &statement
}
