package parser_test

import (
	"taulang/ast"
	"taulang/lexer"
	"taulang/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		expectedErrors  []string
		expectedProgram *ast.Program
	}{
		{
			name:           "success",
			input:          `let x = 5;`,
			expectedErrors: []string{},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l, err := lexer.NewLexer(tc.input)
			assert.NoError(t, err)

			p := parser.NewParser(l)
			program := p.Parse()

			assert.Equal(t, tc.expectedProgram, program)
			assert.Equal(t, tc.expectedErrors, p.Errors())
		})
	}
}
