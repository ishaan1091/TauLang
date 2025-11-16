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
	p.prefixParseFunctions[token.LEFT_PAREN] = p.parseGroupedExpressions
	p.prefixParseFunctions[token.FUNCTION] = p.parseFunctionLiteral
	p.prefixParseFunctions[token.IF] = p.parseConditionalExpression
	p.prefixParseFunctions[token.WHILE] = p.parseWhileLoop
	p.prefixParseFunctions[token.BREAK] = p.parseBreak
	p.prefixParseFunctions[token.CONTINUE] = p.parseContinue

	p.infixParseFunctions[token.EQUALS] = p.parseInfixExpression
	p.infixParseFunctions[token.NOT_EQUALS] = p.parseInfixExpression
	p.infixParseFunctions[token.ADDITION] = p.parseInfixExpression
	p.infixParseFunctions[token.SUBTRACTION] = p.parseInfixExpression
	p.infixParseFunctions[token.DIVISION] = p.parseInfixExpression
	p.infixParseFunctions[token.MULTIPLICATION] = p.parseInfixExpression
	p.infixParseFunctions[token.LESSER_THAN] = p.parseInfixExpression
	p.infixParseFunctions[token.LESSER_EQUALS] = p.parseInfixExpression
	p.infixParseFunctions[token.GREATER_THAN] = p.parseInfixExpression
	p.infixParseFunctions[token.GREATER_EQUALS] = p.parseInfixExpression

	p.infixParseFunctions[token.LEFT_PAREN] = p.parseCallExpression

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

func (p *parser) noInfixParseFunctionError(tok token.Token) {
	msg := fmt.Sprintf("no infix parse function found for %s", tok.Type)
	if tok.Type == token.ILLEGAL {
		msg += fmt.Sprintf(" (%s)", tok.Literal)
	}
	p.errors = append(p.errors, msg)
}

func (p *parser) callExpressionPeekTokenMismatchError() {
	p.errors = append(p.errors, fmt.Sprintf("expected next token to be , or ) but got %s", p.peekToken.Literal))
}

func (p *parser) parseStatement() ast.Statement {
	// TODO: Implement Function statement so that we can define functions without using
	// 		 let statements as well, like we do in golang
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.IDENTIFIER:
		// Lookahead: if identifier followed by '=' that means it is assignment statement
		if p.peekTokenIs(token.ASSIGNMENT) {
			return p.parseAssignmentStatement()
		}
		return p.parseExpressionStatement()
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

func (p *parser) parseAssignmentStatement() ast.Statement {
	statement := ast.AssignmentStatement{Token: p.currToken}

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

	// Here breaking at semicolon would have also automatically been handled
	// by precedence check (since semicolon has no precedence hence gets defauls
	// precedence of LOWEST, which is the smallest value hence it can't be greater)
	// But we kept this check explicit for better readability
	for !p.peekTokenIs(token.SEMICOLON) && p.peekPrecedence() > precedence {
		infixParser := p.infixParseFunctions[p.peekToken.Type]
		if infixParser == nil {
			p.noInfixParseFunctionError(p.peekToken)
			return nil
		}

		p.nextToken()

		// to understand why we did so dry run this code on following example
		// eg - 1 * 2 + 3
		left = infixParser(left)
	}

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
	val, err := strconv.ParseInt(p.currToken.Literal, 0, 64)
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

func (p *parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := ast.InfixExpression{Token: p.currToken, Left: left, Operator: p.currToken.Literal}

	precedence := p.currPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return &expression
}

func (p *parser) parseCallExpression(function ast.Expression) ast.Expression {
	expression := ast.CallExpression{Token: p.currToken, Function: function}

	if p.peekTokenIs(token.RIGHT_PAREN) {
		p.nextToken()
		expression.Arguments = []ast.Expression{}
		return &expression
	}

	var args []ast.Expression
	for !p.currTokenIs(token.RIGHT_PAREN) {
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))

		if !p.peekTokenIs(token.COMMA) && !p.peekTokenIs(token.RIGHT_PAREN) {
			p.callExpressionPeekTokenMismatchError()
			return nil
		}
		p.nextToken()
	}

	expression.Arguments = args

	return &expression
}

func (p *parser) parseGroupedExpressions() ast.Expression {
	p.nextToken()

	expression := p.parseExpression(LOWEST)

	if !p.expectPeekToken(token.RIGHT_PAREN) {
		return nil
	}

	// Here we don't need a separate node for Grouped expression as there purpose
	// which was evaluating expression within parenthesis first got solved as we are
	// forcing such behavior and hence these parenthesis have no significance anymore
	// To better understand dry run following example - 1 + (2 + 3) + 4
	return expression
}

func (p *parser) parseFunctionLiteral() ast.Expression {
	expression := ast.FunctionLiteral{Token: p.currToken}

	if !p.expectPeekToken(token.LEFT_PAREN) {
		return nil
	}

	var params []*ast.Identifier
	for !p.currTokenIs(token.RIGHT_PAREN) {
		p.nextToken()
		param := p.parseExpression(LOWEST)
		ident, ok := param.(*ast.Identifier)
		if !ok {
			p.errors = append(p.errors, fmt.Sprintf("expected IDENTIFIER in function parameters got: %s", param.String()))
			return nil
		}
		params = append(params, ident)

		if !p.peekTokenIs(token.COMMA) && !p.peekTokenIs(token.RIGHT_PAREN) {
			p.callExpressionPeekTokenMismatchError()
			return nil
		}
		p.nextToken()
	}
	expression.Parameters = params

	if !p.expectPeekToken(token.LEFT_BRACE) {
		return nil
	}

	expression.Body = p.parseBlockStatement()

	return &expression
}

func (p *parser) parseBlockStatement() *ast.BlockStatement {
	block := ast.BlockStatement{Token: p.currToken}

	p.nextToken()

	var statements []ast.Statement
	for !p.currTokenIs(token.RIGHT_BRACE) && !p.currTokenIs(token.EOF) {
		statements = append(statements, p.parseStatement())
		p.nextToken()
	}

	if p.currTokenIs(token.EOF) {
		p.errors = append(p.errors, fmt.Sprintf("expected next token to be RIGHT_BRACE, found EOF"))
		return nil
	}

	block.Statements = statements

	return &block
}

func (p *parser) parseConditionalExpression() ast.Expression {
	expression := ast.ConditionalExpression{Token: p.currToken}

	if !p.expectPeekToken(token.LEFT_PAREN) {
		return nil
	}

	p.nextToken()
	expression.Condition = p.parseExpression(LOWEST)

	if !p.expectPeekToken(token.RIGHT_PAREN) {
		return nil
	}

	if !p.expectPeekToken(token.LEFT_BRACE) {
		return nil
	}

	expression.Consequence = p.parseBlockStatement()

	if !p.peekTokenIs(token.ELSE) {
		return &expression
	}

	p.nextToken()
	if !p.expectPeekToken(token.LEFT_BRACE) {
		return nil
	}

	expression.Alternative = p.parseBlockStatement()

	return &expression
}

func (p *parser) parseWhileLoop() ast.Expression {
	expression := ast.WhileLoopExpression{Token: p.currToken}

	if !p.expectPeekToken(token.LEFT_PAREN) {
		return nil
	}

	p.nextToken()
	expression.Condition = p.parseExpression(LOWEST)

	if !p.expectPeekToken(token.RIGHT_PAREN) {
		return nil
	}

	if !p.expectPeekToken(token.LEFT_BRACE) {
		return nil
	}

	expression.Body = p.parseBlockStatement()

	return &expression
}

func (p *parser) parseBreak() ast.Expression {
	return &ast.BreakExpression{Token: p.currToken}
}

func (p *parser) parseContinue() ast.Expression {
	return &ast.ContinueExpression{Token: p.currToken}
}
