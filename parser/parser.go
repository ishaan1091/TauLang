package parser

import (
	"fmt"
	"strconv"
	"taulang/ast"
	"taulang/lexer"
	"taulang/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

var precedences = map[token.Type]int{
	token.EQUALS:         EQUALS,
	token.NOT_EQUALS:     EQUALS,
	token.LESSER_THAN:    LESSGREATER,
	token.LESSER_EQUALS:  LESSGREATER,
	token.GREATER_THAN:   LESSGREATER,
	token.GREATER_EQUALS: LESSGREATER,
	token.ADDITION:       SUM,
	token.SUBTRACTION:    SUM,
	token.DIVISION:       PRODUCT,
	token.MULTIPLICATION: PRODUCT,
	token.LEFT_PAREN:     CALL,
}

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

	p.prefixParseFunctions[token.IDENTIFIER] = p.parseIdentifier
	p.prefixParseFunctions[token.NUMBER] = p.parseIntegerLiteral
	p.prefixParseFunctions[token.BANG] = p.parsePrefixExpression
	p.prefixParseFunctions[token.SUBTRACTION] = p.parsePrefixExpression
	p.prefixParseFunctions[token.TRUE] = p.parseBoolean
	p.prefixParseFunctions[token.FALSE] = p.parseBoolean
	p.prefixParseFunctions[token.STRING] = p.parseString

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

func (p *parser) noPrefixParseFunctionError(tok token.Token) {
	msg := fmt.Sprintf("no prefix parse function found for %s", tok.Type)
	if tok.Type == token.ILLEGAL {
		msg += fmt.Sprintf(" (%s)", tok.Literal)
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

	p.nextToken()
	statement.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return &statement
}

func (p *parser) parseReturnStatement() ast.Statement {
	statement := ast.ReturnStatement{Token: p.currToken}

	p.nextToken()
	statement.ReturnValue = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return &statement
}

func (p *parser) parseExpressionStatement() ast.Statement {
	statement := ast.ExpressionStatement{Token: p.currToken}

	statement.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return &statement
}

func (p *parser) parseExpression(precedence int) ast.Expression {
	prefixParser := p.prefixParseFunctions[p.currToken.Type]
	if prefixParser == nil {
		p.noPrefixParseFunctionError(p.currToken)
		return nil
	}

	left := prefixParser()

	return left
}

func getPrecedence(tok token.Type) int {
	if precedence, ok := precedences[tok]; ok {
		return precedence
	}

	return LOWEST
}

func (p *parser) currPrecedence() int {
	return getPrecedence(p.currToken.Type)
}

func (p *parser) peekPrecedence() int {
	return getPrecedence(p.peekToken.Type)
}

func (p *parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
}

func (p *parser) parseIntegerLiteral() ast.Expression {
	expression := ast.IntegerLiteral{Token: p.currToken}

	// TODO: Add support to parse decimal values
	val, err := strconv.Atoi(p.currToken.Literal)
	if err != nil {
		p.errors = append(p.errors, fmt.Sprintf("could not parse %q as integer", p.currToken.Literal))
		return nil
	}
	expression.Value = val

	return &expression
}

func (p *parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.currToken, Value: p.currTokenIs(token.TRUE)}
}

func (p *parser) parseString() ast.Expression {
	return &ast.String{Token: p.currToken, Value: p.currToken.Literal}
}

func (p *parser) parsePrefixExpression() ast.Expression {
	expression := ast.PrefixExpression{Token: p.currToken, Operator: p.currToken.Literal}

	p.nextToken()

	expression.Operand = p.parseExpression(PREFIX)

	return &expression
}
