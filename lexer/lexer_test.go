package lexer

import (
	"fmt"
	"taulang/token"
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
			l, err := NewLexer(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, l)
		})
	}
}

func TestLexer(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected token.Token
	}{
		{
			name:  "left brace",
			input: "{",
			expected: token.Token{
				Type:    token.LEFT_BRACE,
				Literal: "{",
			},
		},
		{
			name:  "right brace",
			input: "}",
			expected: token.Token{
				Type:    token.RIGHT_BRACE,
				Literal: "}",
			},
		},
		{
			name:  "left bracket",
			input: "[",
			expected: token.Token{
				Type:    token.LEFT_BRACKET,
				Literal: "[",
			},
		},
		{
			name:  "right bracket",
			input: "]",
			expected: token.Token{
				Type:    token.RIGHT_BRACKET,
				Literal: "]",
			},
		},
		{
			name:  "colon",
			input: ":",
			expected: token.Token{
				Type:    token.COLON,
				Literal: ":",
			},
		},
		{
			name:  "comma",
			input: ",",
			expected: token.Token{
				Type:    token.COMMA,
				Literal: ",",
			},
		},
		{
			name:  "semicolon",
			input: ";",
			expected: token.Token{
				Type:    token.SEMICOLON,
				Literal: ";",
			},
		},
		{
			name:  "equals",
			input: "==",
			expected: token.Token{
				Type:    token.EQUALS,
				Literal: "==",
			},
		},
		{
			name:  "not equals",
			input: "!=",
			expected: token.Token{
				Type:    token.NOT_EQUALS,
				Literal: "!=",
			},
		},
		{
			name:  "greater than",
			input: ">",
			expected: token.Token{
				Type:    token.GREATER_THAN,
				Literal: ">",
			},
		},
		{
			name:  "less than",
			input: "<",
			expected: token.Token{
				Type:    token.LESSER_THAN,
				Literal: "<",
			},
		},
		{
			name:  "addition",
			input: "+",
			expected: token.Token{
				Type:    token.ADDITION,
				Literal: "+",
			},
		},
		{
			name:  "subtraction",
			input: "-",
			expected: token.Token{
				Type:    token.SUBTRACTION,
				Literal: "-",
			},
		},
		{
			name:  "multiplication",
			input: "*",
			expected: token.Token{
				Type:    token.MULTIPLICATION,
				Literal: "*",
			},
		},
		{
			name:  "division",
			input: "/",
			expected: token.Token{
				Type:    token.DIVISION,
				Literal: "/",
			},
		},
		{
			name:  "string",
			input: "\"test\"",
			expected: token.Token{
				Type:    token.STRING,
				Literal: "test",
			},
		},
		{
			name:  "number",
			input: "123",
			expected: token.Token{
				Type:    token.NUMBER,
				Literal: "123",
			},
		},
		{
			name:  "illegal",
			input: "@",
			expected: token.Token{
				Type:    token.ILLEGAL,
				Literal: "@",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, err := NewLexer(tt.input)
			assert.NoError(t, err)
			tok := l.NextToken()
			assert.Equal(t, tt.expected, tok)
		})
	}
}

func TestLexerMultipleTokens(t *testing.T) {
	input := `
	{
		"key": "value",
		"number": 123,
		"array": [1, 2, 3],
		"object": {"nestedKey": "nestedValue"},
		"boolean": true,
		"comparison": 5 > 3,
		"arithmetic": 10 + 20 * 3
		"nested": {"key": "value"}
	}


	sun_liyo_tau x ne_bana_diye 5;
	agar_maan_lo (x > 0) {
		laadle_ye_le x;
	} na_toh {
		laadle_ye_le 0;
	}

	sun_liyo_tau add ne_bana_diye rasoi_mein_bata_diye(y, z) {
		laadle_ye_le y + z;
	};

	jab_tak (x < 10) {
		x ne_bana_diye x + 1;
	}

	rok_diye;
	jaan_de;

	sun_liyo_tau arr ne_bana_diye [1, 2, 3, 4, 5];
	sun_liyo_tau obj ne_bana_diye {"a": 1, "b": 2, "c": 3};
	sun_liyo_tau bool ne_bana_diye saccha;
	sun_liyo_tau bool ne_bana_diye jhootha;
	`

	expectedTokens := []token.Token{
		{Type: token.LEFT_BRACE, Literal: "{"},
		{Type: token.STRING, Literal: "key"},
		{Type: token.COLON, Literal: ":"},
		{Type: token.STRING, Literal: "value"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.STRING, Literal: "number"},
		{Type: token.COLON, Literal: ":"},
		{Type: token.NUMBER, Literal: "123"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.STRING, Literal: "array"},
		{Type: token.COLON, Literal: ":"},
		{Type: token.LEFT_BRACKET, Literal: "["},
		{Type: token.NUMBER, Literal: "1"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.NUMBER, Literal: "2"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.NUMBER, Literal: "3"},
		{Type: token.RIGHT_BRACKET, Literal: "]"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.STRING, Literal: "object"},
		{Type: token.COLON, Literal: ":"},
		{Type: token.LEFT_BRACE, Literal: "{"},
		{Type: token.STRING, Literal: "nestedKey"},
		{Type: token.COLON, Literal: ":"},
		{Type: token.STRING, Literal: "nestedValue"},
		{Type: token.RIGHT_BRACE, Literal: "}"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.STRING, Literal: "boolean"},
		{Type: token.COLON, Literal: ":"},
		{Type: token.IDENTIFIER, Literal: "true"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.STRING, Literal: "comparison"},
		{Type: token.COLON, Literal: ":"},
		{Type: token.NUMBER, Literal: "5"},
		{Type: token.GREATER_THAN, Literal: ">"},
		{Type: token.NUMBER, Literal: "3"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.STRING, Literal: "arithmetic"},
		{Type: token.COLON, Literal: ":"},
		{Type: token.NUMBER, Literal: "10"},
		{Type: token.ADDITION, Literal: "+"},
		{Type: token.NUMBER, Literal: "20"},
		{Type: token.MULTIPLICATION, Literal: "*"},
		{Type: token.NUMBER, Literal: "3"},
		{Type: token.STRING, Literal: "nested"},
		{Type: token.COLON, Literal: ":"},
		{Type: token.LEFT_BRACE, Literal: "{"},
		{Type: token.STRING, Literal: "key"},
		{Type: token.COLON, Literal: ":"},
		{Type: token.STRING, Literal: "value"},
		{Type: token.RIGHT_BRACE, Literal: "}"},
		{Type: token.RIGHT_BRACE, Literal: "}"},
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "x"},
		{Type: token.ASSIGNMENT, Literal: "="},
		{Type: token.NUMBER, Literal: "5"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.IF, Literal: "if"},
		{Type: token.LEFT_PAREN, Literal: "("},
		{Type: token.IDENTIFIER, Literal: "x"},
		{Type: token.GREATER_THAN, Literal: ">"},
		{Type: token.NUMBER, Literal: "0"},
		{Type: token.RIGHT_PAREN, Literal: ")"},
		{Type: token.LEFT_BRACE, Literal: "{"},
		{Type: token.RETURN, Literal: "return"},
		{Type: token.IDENTIFIER, Literal: "x"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.RIGHT_BRACE, Literal: "}"},
		{Type: token.ELSE, Literal: "else"},
		{Type: token.LEFT_BRACE, Literal: "{"},
		{Type: token.RETURN, Literal: "return"},
		{Type: token.NUMBER, Literal: "0"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.RIGHT_BRACE, Literal: "}"},
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "add"},
		{Type: token.ASSIGNMENT, Literal: "="},
		{Type: token.FUNCTION, Literal: "func"},
		{Type: token.LEFT_PAREN, Literal: "("},
		{Type: token.IDENTIFIER, Literal: "y"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.IDENTIFIER, Literal: "z"},
		{Type: token.RIGHT_PAREN, Literal: ")"},
		{Type: token.LEFT_BRACE, Literal: "{"},
		{Type: token.RETURN, Literal: "return"},
		{Type: token.IDENTIFIER, Literal: "y"},
		{Type: token.ADDITION, Literal: "+"},
		{Type: token.IDENTIFIER, Literal: "z"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.RIGHT_BRACE, Literal: "}"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.WHILE, Literal: "while"},
		{Type: token.LEFT_PAREN, Literal: "("},
		{Type: token.IDENTIFIER, Literal: "x"},
		{Type: token.LESSER_THAN, Literal: "<"},
		{Type: token.NUMBER, Literal: "10"},
		{Type: token.RIGHT_PAREN, Literal: ")"},
		{Type: token.LEFT_BRACE, Literal: "{"},
		{Type: token.IDENTIFIER, Literal: "x"},
		{Type: token.ASSIGNMENT, Literal: "="},
		{Type: token.IDENTIFIER, Literal: "x"},
		{Type: token.ADDITION, Literal: "+"},
		{Type: token.NUMBER, Literal: "1"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.RIGHT_BRACE, Literal: "}"},
		{Type: token.BREAK, Literal: "break"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.CONTINUE, Literal: "continue"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "arr"},
		{Type: token.ASSIGNMENT, Literal: "="},
		{Type: token.LEFT_BRACKET, Literal: "["},
		{Type: token.NUMBER, Literal: "1"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.NUMBER, Literal: "2"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.NUMBER, Literal: "3"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.NUMBER, Literal: "4"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.NUMBER, Literal: "5"},
		{Type: token.RIGHT_BRACKET, Literal: "]"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "obj"},
		{Type: token.ASSIGNMENT, Literal: "="},
		{Type: token.LEFT_BRACE, Literal: "{"},
		{Type: token.STRING, Literal: "a"},
		{Type: token.COLON, Literal: ":"},
		{Type: token.NUMBER, Literal: "1"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.STRING, Literal: "b"},
		{Type: token.COLON, Literal: ":"},
		{Type: token.NUMBER, Literal: "2"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.STRING, Literal: "c"},
		{Type: token.COLON, Literal: ":"},
		{Type: token.NUMBER, Literal: "3"},
		{Type: token.RIGHT_BRACE, Literal: "}"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "bool"},
		{Type: token.ASSIGNMENT, Literal: "="},
		{Type: token.TRUE, Literal: "true"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.LET, Literal: "let"},
		{Type: token.IDENTIFIER, Literal: "bool"},
		{Type: token.ASSIGNMENT, Literal: "="},
		{Type: token.FALSE, Literal: "false"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	}

	l, err := NewLexer(input)
	assert.NoError(t, err)

	for i, expectedToken := range expectedTokens {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if i == 99 {
				t.Log(expectedToken)
			}
			tok := l.NextToken()
			assert.Equal(t, expectedToken, tok)
		})
	}
}
