package lexer

import (
	"errors"
	"jsonparser/cmd/jsonparser/commons"
	"unicode/utf8"
)

func Tokenize(content string) ([]*commons.Token, error) {
	tokens := []*commons.Token{}

	for i, w := 0, 0; i < len(content); i += w {
		ch, width, err := readNextChar(content, i)
		if err != nil {
			return nil, err
		}

		w = width

		switch ch {
		case '{':
			tokens = append(tokens, &commons.Token{Type: commons.LEFT_BRACE, Literal: "{"})
		case '}':
			tokens = append(tokens, &commons.Token{Type: commons.RIGHT_BRACE, Literal: "}"})
		case '[':
			tokens = append(tokens, &commons.Token{Type: commons.LEFT_BRACKET, Literal: "["})
		case ']':
			tokens = append(tokens, &commons.Token{Type: commons.RIGHT_BRACKET, Literal: "]"})
		case ':':
			tokens = append(tokens, &commons.Token{Type: commons.COLON, Literal: ":"})
		case ',':
			tokens = append(tokens, &commons.Token{Type: commons.COMMA, Literal: ","})
		case '"':
			str, bytesReadWidth, err := readString(content, i)
			if err != nil {
				return nil, err
			}

			w = bytesReadWidth

			tokens = append(tokens, &commons.Token{Type: commons.STRING, Literal: str})
		default:
			// Ignore white spaces
			// Read values until next whitespace
			// Validate that they should be either number (int or float) or boolean
			// Create and append token accordingly
		}
	}

	tokens = append(tokens, &commons.Token{Type: commons.EOF, Literal: ""})

	return tokens, nil
}

func readString(content string, pos int) (string, int, error) {
	bytesReadWidth := 1
	str := ""

	for i, w := pos+1, 0; ; i += w {
		ch, width, err := readNextChar(content, i)
		if err != nil {
			return "", 0, err
		}

		w = width
		bytesReadWidth += width

		if ch == '"' {
			break
		}

		str += string(ch)
	}

	return str, bytesReadWidth, nil
}

func readNextChar(content string, pos int) (rune, int, error) {
	runeValue, width := utf8.DecodeRuneInString(content[pos:])

	if runeValue == utf8.RuneError {
		if width == 0 {
			return 0, 0, errors.New("empty string can't be decoded")
		}
		return 0, 0, errors.New("encoding other than UTF-8 not allowed")
	}

	return runeValue, width, nil
}
