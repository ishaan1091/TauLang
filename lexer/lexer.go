package lexer

import (
	"errors"
	"taulang/commons"
	"unicode"
	"unicode/utf8"
)

const EOF = rune(0)

type Lexer interface {
	NextToken() (commons.Token, error)
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

func (l *lexer) NextToken() (commons.Token, error) {
	var token commons.Token

	switch l.currChar {
	case '{':
		token = commons.NewToken(commons.LEFT_BRACE, "{")
	case '}':
		token = commons.NewToken(commons.RIGHT_BRACE, "}")
	case '[':
		token = commons.NewToken(commons.LEFT_BRACKET, "[")
	case ']':
		token = commons.NewToken(commons.RIGHT_BRACKET, "]")
	case '(':
		token = commons.NewToken(commons.LEFT_PAREN, "(")
	case ')':
		token = commons.NewToken(commons.RIGHT_PAREN, ")")
	case ':':
		token = commons.NewToken(commons.COLON, ":")
	case ',':
		token = commons.NewToken(commons.COMMA, ",")
	case ';':
		token = commons.NewToken(commons.SEMICOLON, ";")
	case '=':
		token = l.readEqualsOrDefaultToken(commons.EQUALS, commons.ILLEGAL)
	case '!':
		token = l.readEqualsOrDefaultToken(commons.NOT_EQUALS, commons.BANG)
	case '>':
		token = l.readEqualsOrDefaultToken(commons.GREATER_EQUALS, commons.GREATER_THAN)
	case '<':
		token = l.readEqualsOrDefaultToken(commons.LESSER_EQUALS, commons.LESSER_THAN)
	case '+':
		token = commons.NewToken(commons.ADDITION, "+")
	case '-':
		token = commons.NewToken(commons.SUBTRACTION, "-")
	case '*':
		token = commons.NewToken(commons.MULTIPLICATION, "*")
	case '/':
		token = commons.NewToken(commons.DIVISION, "/")
	case '"':
		stringLiteral, err := l.readString()
		if err != nil {
			return commons.Token{}, err
		}
		token = commons.NewToken(commons.STRING, stringLiteral)
	case EOF:
		token = commons.NewToken(commons.EOF, "")
	default:
		if unicode.IsLetter(l.currChar) {
			identifier, err := l.readIdentifier()
			if err != nil {
				return commons.Token{}, err
			}

			token = commons.GetTokenForIdentifierOrKeyword(identifier)
		} else if unicode.IsNumber(l.currChar) {
			number, err := l.readNumber()
			if err != nil {
				return commons.Token{}, err
			}
			token = commons.NewToken(commons.NUMBER, number)
		} else {
			token = commons.NewToken(commons.ILLEGAL, string(l.currChar))
		}
	}

	err := l.readNextChar()
	if err != nil {
		return commons.Token{}, err
	}

	err = l.skipWhiteSpaces()
	if err != nil {
		return commons.Token{}, err
	}

	return token, nil
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

func (l *lexer) readEqualsOrDefaultToken(compoundType commons.TokenType, defaultType commons.TokenType) commons.Token {
	if nextChar, _, err := l.decodeNextChar(); err == nil && nextChar == '=' {
		currChar := l.currChar
		l.readNextChar()
		return commons.NewToken(compoundType, string(currChar)+string(nextChar))
	}
	return commons.NewToken(defaultType, string(l.currChar))
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
