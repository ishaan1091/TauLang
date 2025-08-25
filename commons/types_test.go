package commons

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookupIdentifier(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected TokenType
	}{
		{
			name:     "lookup LET keyword",
			input:    "sun_liyo_tau",
			expected: LET,
		},
		{
			name:     "lookup FUNCTION keyword",
			input:    "rasoi_mein_bata_diye",
			expected: FUNCTION,
		},
		{
			name:     "lookup IF keyword",
			input:    "agar_maan_lo",
			expected: IF,
		},
		{
			name:     "lookup ELSE keyword",
			input:    "na_toh",
			expected: ELSE,
		},
		{
			name:     "lookup RETURN keyword",
			input:    "laadle_ye_le",
			expected: RETURN,
		},
		{
			name:     "lookup TRUE keyword",
			input:    "saccha",
			expected: TRUE,
		},
		{
			name:     "lookup FALSE keyword",
			input:    "jhootha",
			expected: FALSE,
		},
		{
			name:     "lookup WHILE keyword",
			input:    "jab_tak",
			expected: WHILE,
		},
		{
			name:     "lookup BREAK keyword",
			input:    "rok_diye",
			expected: BREAK,
		},
		{
			name:     "lookup CONTINUE keyword",
			input:    "jaan_de",
			expected: CONTINUE,
		},
		{
			name:     "lookup identifier",
			input:    "myVariable",
			expected: IDENTIFIER,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := LookupIdentifier(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
