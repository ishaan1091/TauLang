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
			name:           "failure - illegal token",
			input:          `sun_liyo_tau x = 5;`,
			expectedErrors: []string{"expected next token to be ASSIGNMENT, got ILLEGAL (=)"},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{},
			},
		},
		{
			name:           "success - parse let statement",
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
			name:           "success - parse let statement with boolean value",
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
			name:           "success - parse let statement with identifier value",
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
		{
			name:           "failure - parse let statement with missing identifier",
			input:          `sun_liyo_tau ne_bana_diye x y;`,
			expectedErrors: []string{"expected next token to be IDENTIFIER, got ASSIGNMENT"},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{},
			},
		},
		{
			name:           "success - return statement",
			input:          `laadle_ye_le x;`,
			expectedErrors: []string{},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.ReturnStatement{
						Token:       token.Token{Type: token.RETURN, Literal: "return"},
						ReturnValue: nil,
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
