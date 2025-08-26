package lexer

import (
	"taulang/commons"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLexer(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *lexer
	}{
		{
			name:  "empty input",
			input: "",
			expected: &lexer{
				source:           "",
				currCharPosition: 0,
				nextCharPosition: 0,
				currChar:         0,
			},
		},
		{
			name:  "multiple characters",
			input: "abc",
			expected: &lexer{
				source:           "abc",
				currCharPosition: 0,
				nextCharPosition: 1,
				currChar:         'a',
			},
		},
		{
			name:  "escaped string",
			input: "\\\"test\\\"",
			expected: &lexer{
				source:           "\\\"test\\\"",
				currCharPosition: 0,
				nextCharPosition: 1,
				currChar:         '\\',
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.input)
			assert.Equal(t, tt.expected, l)
		})
	}
}

func TestLexer(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected commons.Token
	}{
		{
			name:  "left brace",
			input: "{",
			expected: commons.Token{
				Type:    commons.LEFT_BRACE,
				Literal: "{",
			},
		},
		{
			name:  "right brace",
			input: "}",
			expected: commons.Token{
				Type:    commons.RIGHT_BRACE,
				Literal: "}",
			},
		},
		{
			name:  "left bracket",
			input: "[",
			expected: commons.Token{
				Type:    commons.LEFT_BRACKET,
				Literal: "[",
			},
		},
		{
			name:  "right bracket",
			input: "]",
			expected: commons.Token{
				Type:    commons.RIGHT_BRACKET,
				Literal: "]",
			},
		},
		{
			name:  "colon",
			input: ":",
			expected: commons.Token{
				Type:    commons.COLON,
				Literal: ":",
			},
		},
		{
			name:  "comma",
			input: ",",
			expected: commons.Token{
				Type:    commons.COMMA,
				Literal: ",",
			},
		},
		{
			name:  "semicolon",
			input: ";",
			expected: commons.Token{
				Type:    commons.SEMICOLON,
				Literal: ";",
			},
		},
		{
			name:  "equals",
			input: "==",
			expected: commons.Token{
				Type:    commons.EQUALS,
				Literal: "==",
			},
		},
		{
			name:  "not equals",
			input: "!=",
			expected: commons.Token{
				Type:    commons.NOT_EQUALS,
				Literal: "!=",
			},
		},
		{
			name:  "greater than",
			input: ">",
			expected: commons.Token{
				Type:    commons.GREATER_THAN,
				Literal: ">",
			},
		},
		{
			name:  "less than",
			input: "<",
			expected: commons.Token{
				Type:    commons.LESSER_THAN,
				Literal: "<",
			},
		},
		{
			name:  "addition",
			input: "+",
			expected: commons.Token{
				Type:    commons.ADDITION,
				Literal: "+",
			},
		},
		{
			name:  "subtraction",
			input: "-",
			expected: commons.Token{
				Type:    commons.SUBTRACTION,
				Literal: "-",
			},
		},
		{
			name:  "multiplication",
			input: "*",
			expected: commons.Token{
				Type:    commons.MULTIPLICATION,
				Literal: "*",
			},
		},
		{
			name:  "division",
			input: "/",
			expected: commons.Token{
				Type:    commons.DIVISION,
				Literal: "/",
			},
		},
		{
			name:  "string",
			input: "\"test\"",
			expected: commons.Token{
				Type:    commons.STRING,
				Literal: "test",
			},
		},
		{
			name:  "number",
			input: "123",
			expected: commons.Token{
				Type:    commons.NUMBER,
				Literal: "123",
			},
		},
		{
			name:  "illegal",
			input: "@",
			expected: commons.Token{
				Type:    commons.ILLEGAL,
				Literal: "@",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.input)
			token, err := l.NextToken()
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, token)
		})
	}
}
