package lexer

import (
	"errors"
	"taulang/commons"
	"unicode/utf8"
)

type Lexer interface {
	NextToken() (commons.Token, error)
	Peek() (commons.Token, error)
}

type lexer struct {
	source           string
	currCharPosition int
	nextCharPosition int
	currChar         rune
}

func NewLexer(input string) *lexer {
	l := lexer{
		source:           input,
		currCharPosition: 0,
		nextCharPosition: 0,
		currChar:         0,
	}
	return &l
}

func (l *lexer) NextToken() (commons.Token, error) {
	var token commons.Token

	switch l.currChar {
	case '{':
		token = newToken(commons.LEFT_BRACE, "{")
	case '}':
		token = newToken(commons.RIGHT_BRACE, "}")
	case '[':
		token = newToken(commons.LEFT_BRACKET, "[")
	case ']':
		token = newToken(commons.RIGHT_BRACKET, "]")
	case ':':
		token = newToken(commons.COLON, ":")
	case ',':
		token = newToken(commons.COMMA, ",")
	case ';':
		token = newToken(commons.SEMICOLON, ";")
	case '=':
		token = l.readEqualsOrDefaultToken(commons.EQUALS, commons.ILLEGAL)
	case '!':
		token = l.readEqualsOrDefaultToken(commons.NOT_EQUALS, commons.BANG)
	case '>':
		token = l.readEqualsOrDefaultToken(commons.GREATER_EQUALS, commons.GREATER_THAN)
	case '<':
		token = l.readEqualsOrDefaultToken(commons.LESSER_EQUALS, commons.LESSER_THAN)
	case '+':
		token = newToken(commons.ADDITION, "+")
	case '-':
		token = newToken(commons.SUBTRACTION, "-")
	case '*':
		token = newToken(commons.MULTIPLICATION, "*")
	case '/':
		token = newToken(commons.DIVISION, "/")
	case '"':
		// token = newToken(commons.STRING, l.readString())
		token = newToken(commons.ILLEGAL, "")
	case 0:
		token = newToken(commons.EOF, "")
	default:
		// handle tokenising identifiers and keywords
		// handle boolean
		// handle numbers
		token = newToken(commons.ILLEGAL, "")
	}

	l.readNextChar()
	l.skipWhiteSpaces()

	return token, nil
}

func (l *lexer) Peek() (commons.Token, error) {
	return commons.Token{}, nil
}

func (l *lexer) readNextChar() error {
	return nil
}

func (l *lexer) skipWhiteSpaces() error {
	for l.currChar == ' ' || l.currChar == '\t' || l.currChar == '\n' || l.currChar == '\r' {
		if err := l.readNextChar(); err != nil {
			return err
		}
	}
	return nil
}

func (l *lexer) readEqualsOrDefaultToken(compoundType commons.TokenType, defaultType commons.TokenType) commons.Token {
	if nextChar, _, err := l.decodeNextChar(); err == nil && nextChar == '=' {
		return newToken(compoundType, string(l.currChar)+string(nextChar))
	}
	return newToken(defaultType, string(l.currChar))
}

func newToken(tokenType commons.TokenType, literal string) commons.Token {
	return commons.Token{Type: tokenType, Literal: literal}
}

func (l *lexer) decodeNextChar() (rune, int, error) {
	runeValue, width := utf8.DecodeRuneInString(l.source[l.nextCharPosition:])

	if runeValue == utf8.RuneError {
		if width == 0 {
			return 0, 0, errors.New("empty string can't be decoded")
		}
		return 0, 0, errors.New("encoding other than UTF-8 not allowed")
	}

	return runeValue, width, nil
}
