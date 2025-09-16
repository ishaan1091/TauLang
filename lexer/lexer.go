package lexer

import (
	"errors"
	"taulang/token"
	"unicode"
	"unicode/utf8"
)

const EOF = rune(0)

type Lexer interface {
	NextToken() (token.Token, error)
}

type lexer struct {
	source           string
	currCharPosition int
	nextCharPosition int
	currChar         rune
}

func NewLexer(input string) Lexer {
	l := lexer{
		source:           input,
		currCharPosition: 0,
		nextCharPosition: 0,
		currChar:         0,
	}
	l.readNextChar()
	l.skipWhiteSpaces()
	return &l
}

func (l *lexer) NextToken() (token.Token, error) {
	var tok token.Token

	switch l.currChar {
	case '{':
		tok = token.NewToken(token.LEFT_BRACE, "{")
	case '}':
		tok = token.NewToken(token.RIGHT_BRACE, "}")
	case '[':
		tok = token.NewToken(token.LEFT_BRACKET, "[")
	case ']':
		tok = token.NewToken(token.RIGHT_BRACKET, "]")
	case '(':
		tok = token.NewToken(token.LEFT_PAREN, "(")
	case ')':
		tok = token.NewToken(token.RIGHT_PAREN, ")")
	case ':':
		tok = token.NewToken(token.COLON, ":")
	case ',':
		tok = token.NewToken(token.COMMA, ",")
	case ';':
		tok = token.NewToken(token.SEMICOLON, ";")
	case '=':
		tok = l.readEqualsOrDefaultToken(token.EQUALS, token.ILLEGAL)
	case '!':
		tok = l.readEqualsOrDefaultToken(token.NOT_EQUALS, token.BANG)
	case '>':
		tok = l.readEqualsOrDefaultToken(token.GREATER_EQUALS, token.GREATER_THAN)
	case '<':
		tok = l.readEqualsOrDefaultToken(token.LESSER_EQUALS, token.LESSER_THAN)
	case '+':
		tok = token.NewToken(token.ADDITION, "+")
	case '-':
		tok = token.NewToken(token.SUBTRACTION, "-")
	case '*':
		tok = token.NewToken(token.MULTIPLICATION, "*")
	case '/':
		tok = token.NewToken(token.DIVISION, "/")
	case '"':
		stringLiteral, err := l.readString()
		if err != nil {
			return token.Token{}, err
		}
		tok = token.NewToken(token.STRING, stringLiteral)
	case EOF:
		tok = token.NewToken(token.EOF, "")
	default:
		if unicode.IsLetter(l.currChar) {
			identifier, err := l.readIdentifier()
			if err != nil {
				return token.Token{}, err
			}

			tok = token.GetTokenForIdentifierOrKeyword(identifier)
		} else if unicode.IsNumber(l.currChar) {
			number, err := l.readNumber()
			if err != nil {
				return token.Token{}, err
			}
			tok = token.NewToken(token.NUMBER, number)
		} else {
			tok = token.NewToken(token.ILLEGAL, string(l.currChar))
		}
	}

	err := l.readNextChar()
	if err != nil {
		return token.Token{}, err
	}

	err = l.skipWhiteSpaces()
	if err != nil {
		return token.Token{}, err
	}

	return tok, nil
}

func (l *lexer) readNextChar() error {
	if l.nextCharPosition >= len(l.source) {
		l.currChar = EOF
		return nil
	}
	runeValue, width, err := l.decodeNextChar()
	if err != nil {
		return err
	}
	l.currChar = runeValue
	l.currCharPosition = l.nextCharPosition
	l.nextCharPosition += width
	return nil
}

func (l *lexer) skipWhiteSpaces() error {
	for unicode.IsSpace(l.currChar) {
		if err := l.readNextChar(); err != nil {
			return err
		}
	}
	return nil
}

func (l *lexer) readIdentifier() (string, error) {
	var identifier []rune
	for unicode.IsLetter(l.currChar) || unicode.IsNumber(l.currChar) || l.currChar == '_' {
		identifier = append(identifier, l.currChar)

		if nextChar, _, err := l.decodeNextChar(); err == nil && !unicode.IsLetter(nextChar) && !unicode.IsNumber(nextChar) && nextChar != '_' {
			break
		}

		if err := l.readNextChar(); err != nil {
			return "", err
		}
	}
	return string(identifier), nil
}

func (l *lexer) readNumber() (string, error) {
	var number []rune
	dotSeen := false
	for unicode.IsNumber(l.currChar) || l.currChar == '.' {
		number = append(number, l.currChar)

		if l.currChar == '.' {
			if dotSeen {
				return "", errors.New("invalid number")
			}
			dotSeen = true
		}

		if nextChar, _, err := l.decodeNextChar(); err == nil && !unicode.IsNumber(nextChar) && nextChar != '.' {
			break
		}

		if err := l.readNextChar(); err != nil {
			return "", err
		}
	}
	return string(number), nil
}

func (l *lexer) readString() (string, error) {
	var str []rune
	escaped := false
	for {
		if err := l.readNextChar(); err != nil {
			return "", err
		}
		if l.currChar == '"' && !escaped {
			break
		}
		str = append(str, l.currChar)
		escaped = (l.currChar == '\\')
	}
	return string(str), nil
}

func (l *lexer) readEqualsOrDefaultToken(compoundType token.TokenType, defaultType token.TokenType) token.Token {
	if nextChar, _, err := l.decodeNextChar(); err == nil && nextChar == '=' {
		currChar := l.currChar
		l.readNextChar()
		return token.NewToken(compoundType, string(currChar)+string(nextChar))
	}
	return token.NewToken(defaultType, string(l.currChar))
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
