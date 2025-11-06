package parser_test

import (
	"taulang/ast"
	"taulang/lexer"
	"taulang/parser"
	"taulang/token"
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
			name:           "success - parse let simple statement",
			input:          `sun_liyo_tau x ne_bana_diye 5;`,
			expectedErrors: []string{},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.LetStatement{
						Token: token.Token{Type: token.LET, Literal: "let"},
						Name:  &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "x"}, Value: "x"},
						Value: nil,
					},
				},
			},
		},
		{
			name:           "success - parse let simple statement",
			input:          `sun_liyo_tau x ne_bana_diye saccha;`,
			expectedErrors: []string{},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.LetStatement{
						Token: token.Token{Type: token.LET, Literal: "let"},
						Name:  &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "x"}, Value: "x"},
						Value: nil,
					},
				},
			},
		},
		{
			name:           "success - parse let simple statement",
			input:          `sun_liyo_tau x ne_bana_diye y;`,
			expectedErrors: []string{},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.LetStatement{
						Token: token.Token{Type: token.LET, Literal: "let"},
						Name:  &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "x"}, Value: "x"},
						Value: nil,
					},
				},
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
