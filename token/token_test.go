package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookupIdentifier(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		expected        TokenType
		expectedLiteral string
	}{
		{
			name:            "lookup LET keyword",
			input:           "sun_liyo_tau",
			expected:        LET,
			expectedLiteral: "let",
		},
		{
			name:            "lookup FUNCTION keyword",
			input:           "rasoi_mein_bata_diye",
			expected:        FUNCTION,
			expectedLiteral: "func",
		},
		{
			name:            "lookup IF keyword",
			input:           "agar_maan_lo",
			expected:        IF,
			expectedLiteral: "if",
		},
		{
			name:            "lookup ELSE keyword",
			input:           "na_toh",
			expected:        ELSE,
			expectedLiteral: "else",
		},
		{
			name:            "lookup RETURN keyword",
			input:           "laadle_ye_le",
			expected:        RETURN,
			expectedLiteral: "return",
		},
		{
			name:            "lookup TRUE keyword",
			input:           "saccha",
			expected:        TRUE,
			expectedLiteral: "true",
		},
		{
			name:            "lookup FALSE keyword",
			input:           "jhootha",
			expected:        FALSE,
			expectedLiteral: "false",
		},
		{
			name:            "lookup WHILE keyword",
			input:           "jab_tak",
			expected:        WHILE,
			expectedLiteral: "while",
		},
		{
			name:            "lookup BREAK keyword",
			input:           "rok_diye",
			expected:        BREAK,
			expectedLiteral: "break",
		},
		{
			name:            "lookup CONTINUE keyword",
			input:           "jaan_de",
			expected:        CONTINUE,
			expectedLiteral: "continue",
		},
		{
			name:            "lookup identifier",
			input:           "myVariable",
			expected:        IDENTIFIER,
			expectedLiteral: "myVariable",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := GetTokenForIdentifierOrKeyword(tt.input)
			assert.Equal(t, tt.expected, result.Type)
			assert.Equal(t, tt.expectedLiteral, result.Literal)
		})
	}
}
