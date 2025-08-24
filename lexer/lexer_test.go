package lexer

import (
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.input)
			assert.Equal(t, tt.expected, l)
		})
	}
}
