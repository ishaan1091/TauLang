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
			name:  "failure - illegal token",
			input: `sun_liyo_tau x = 5;`,
			expectedErrors: []string{
				"expected next token to be ASSIGNMENT, got ILLEGAL (=)",
				"no prefix parse function found for ILLEGAL (=)",
			},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.ILLEGAL, Literal: "="},
						Expression: nil,
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.NUMBER, Literal: "5"},
						Expression: &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "5"}, Value: 5},
					},
				},
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
						Value: &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "5"}, Value: 5},
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
						Value: &ast.Boolean{Token: token.Token{Type: token.TRUE, Literal: "true"}, Value: true},
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
						Value: &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "y"}, Value: "y"},
					},
				},
			},
		},
		{
			name:  "failure - parse let statement with missing identifier",
			input: `sun_liyo_tau ne_bana_diye x y;`,
			expectedErrors: []string{
				"expected next token to be IDENTIFIER, got ASSIGNMENT",
				"no prefix parse function found for ASSIGNMENT",
			},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.ASSIGNMENT, Literal: "="},
						Expression: nil,
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.IDENTIFIER, Literal: "x"},
						Expression: &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "x"}, Value: "x"},
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.IDENTIFIER, Literal: "y"},
						Expression: &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "y"}, Value: "y"},
					},
				},
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
						ReturnValue: &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "x"}, Value: "x"},
					},
				},
			},
		},
		{
			name:  "success - expression statement - function call",
			input: `some_func(x);`,
			expectedErrors: []string{
				"no prefix parse function found for LEFT_PAREN",
				"no prefix parse function found for RIGHT_PAREN",
			},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.IDENTIFIER, Literal: "some_func"},
						Expression: &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "some_func"}, Value: "some_func"},
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.LEFT_PAREN, Literal: "("},
						Expression: nil,
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.IDENTIFIER, Literal: "x"},
						Expression: &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "x"}, Value: "x"},
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.RIGHT_PAREN, Literal: ")"},
						Expression: nil,
					},
				},
			},
		},
		{
			name: "success - expression statement - function declaration",
			input: `
			rasoi_mein_bata_diye some_func(x) {
				laadle_ye_le x
			};
			`,
			expectedErrors: []string{
				"no prefix parse function found for FUNCTION",
				"no prefix parse function found for LEFT_PAREN",
				"no prefix parse function found for RIGHT_PAREN",
				"no prefix parse function found for LEFT_BRACE",
				"no prefix parse function found for RIGHT_BRACE",
			},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.FUNCTION, Literal: "func"},
						Expression: nil,
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.IDENTIFIER, Literal: "some_func"},
						Expression: &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "some_func"}, Value: "some_func"},
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.LEFT_PAREN, Literal: "("},
						Expression: nil,
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.IDENTIFIER, Literal: "x"},
						Expression: &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "x"}, Value: "x"},
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.RIGHT_PAREN, Literal: ")"},
						Expression: nil,
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.LEFT_BRACE, Literal: "{"},
						Expression: nil,
					},
					&ast.ReturnStatement{
						Token:       token.Token{Type: token.RETURN, Literal: "return"},
						ReturnValue: &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "x"}, Value: "x"},
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.RIGHT_BRACE, Literal: "}"},
						Expression: nil,
					},
				},
			},
		},
		{
			name:           "success - expression statement - standard expression",
			input:          `5 + 5`,
			expectedErrors: []string{"no prefix parse function found for ADDITION"},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.NUMBER, Literal: "5"},
						Expression: &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "5"}, Value: 5},
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.ADDITION, Literal: "+"},
						Expression: nil,
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.NUMBER, Literal: "5"},
						Expression: &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "5"}, Value: 5},
					},
				},
			},
		},
		{
			name:           "success - expression statement - standard expression with identifier",
			input:          `some_var + 5`,
			expectedErrors: []string{"no prefix parse function found for ADDITION"},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.IDENTIFIER, Literal: "some_var"},
						Expression: &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "some_var"}, Value: "some_var"},
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.ADDITION, Literal: "+"},
						Expression: nil,
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.NUMBER, Literal: "5"},
						Expression: &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "5"}, Value: 5},
					},
				},
			},
		},
		{
			name:  "success - expression statement - standard expression with identifier",
			input: `(some_var + 5) * y`,
			expectedErrors: []string{
				"no prefix parse function found for LEFT_PAREN",
				"no prefix parse function found for ADDITION",
				"no prefix parse function found for RIGHT_PAREN",
				"no prefix parse function found for MULTIPLICATION",
			},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.LEFT_PAREN, Literal: "("},
						Expression: nil,
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.IDENTIFIER, Literal: "some_var"},
						Expression: &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "some_var"}, Value: "some_var"},
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.ADDITION, Literal: "+"},
						Expression: nil,
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.NUMBER, Literal: "5"},
						Expression: &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "5"}, Value: 5},
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.RIGHT_PAREN, Literal: ")"},
						Expression: nil,
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.MULTIPLICATION, Literal: "*"},
						Expression: nil,
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.IDENTIFIER, Literal: "y"},
						Expression: &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "y"}, Value: "y"},
					},
				},
			},
		},
		{
			name:  "success - expression statement - standard expression with identifier and prefix operator",
			input: `-5 + some_var + 5`,
			expectedErrors: []string{
				"no prefix parse function found for ADDITION",
				"no prefix parse function found for ADDITION",
			},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: token.Token{Type: token.SUBTRACTION, Literal: "-"},
						Expression: &ast.PrefixExpression{
							Token:    token.Token{Type: token.SUBTRACTION, Literal: "-"},
							Operator: "-",
							Operand:  &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "5"}, Value: 5},
						},
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.ADDITION, Literal: "+"},
						Expression: nil,
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.IDENTIFIER, Literal: "some_var"},
						Expression: &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "some_var"}, Value: "some_var"},
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.ADDITION, Literal: "+"},
						Expression: nil,
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.NUMBER, Literal: "5"},
						Expression: &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "5"}, Value: 5},
					},
				},
			},
		},
		{
			name:  "success - expression statement - standard expression with identifier and bang prefix operator",
			input: `!5 + some_var + 5`,
			expectedErrors: []string{
				"no prefix parse function found for ADDITION",
				"no prefix parse function found for ADDITION",
			},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: token.Token{Type: token.BANG, Literal: "!"},
						Expression: &ast.PrefixExpression{
							Token:    token.Token{Type: token.BANG, Literal: "!"},
							Operator: "!",
							Operand:  &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "5"}, Value: 5},
						},
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.ADDITION, Literal: "+"},
						Expression: nil,
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.IDENTIFIER, Literal: "some_var"},
						Expression: &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "some_var"}, Value: "some_var"},
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.ADDITION, Literal: "+"},
						Expression: nil,
					},
					&ast.ExpressionStatement{
						Token:      token.Token{Type: token.NUMBER, Literal: "5"},
						Expression: &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "5"}, Value: 5},
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
