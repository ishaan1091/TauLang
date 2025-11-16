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
			name:           "success - parse let statement with string value",
			input:          `sun_liyo_tau x ne_bana_diye "test";`,
			expectedErrors: []string{},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.LetStatement{
						Token: token.Token{Type: token.LET, Literal: "let"},
						Name:  &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "x"}, Value: "x"},
						Value: &ast.String{Token: token.Token{Type: token.STRING, Literal: "test"}, Value: "test"},
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
			name:           "success - expression statement - function call 1",
			input:          `some_func(x);`,
			expectedErrors: []string{},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: token.Token{Type: token.IDENTIFIER, Literal: "some_func"},
						Expression: &ast.CallExpression{
							Token:    token.Token{Type: token.LEFT_PAREN, Literal: "("},
							Function: &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "some_func"}, Value: "some_func"},
							Arguments: []ast.Expression{
								&ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "x"}, Value: "x"},
							},
						},
					},
				},
			},
		},
		{
			name:           "success - expression statement - function call 2",
			input:          `some_func();`,
			expectedErrors: []string{},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: token.Token{Type: token.IDENTIFIER, Literal: "some_func"},
						Expression: &ast.CallExpression{
							Token:     token.Token{Type: token.LEFT_PAREN, Literal: "("},
							Function:  &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "some_func"}, Value: "some_func"},
							Arguments: []ast.Expression{},
						},
					},
				},
			},
		},
		{
			name:           "success - expression statement - function call 3",
			input:          `ourFunction(20) + first + second;`,
			expectedErrors: []string{},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: token.Token{Type: token.IDENTIFIER, Literal: "ourFunction"},
						Expression: &ast.InfixExpression{
							Token: token.Token{Type: token.ADDITION, Literal: "+"},
							Left: &ast.InfixExpression{
								Token: token.Token{Type: token.ADDITION, Literal: "+"},
								Left: &ast.CallExpression{
									Token:    token.Token{Type: token.LEFT_PAREN, Literal: "("},
									Function: &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "ourFunction"}, Value: "ourFunction"},
									Arguments: []ast.Expression{
										&ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "20"}, Value: 20},
									},
								},
								Operator: "+",
								Right:    &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "first"}, Value: "first"},
							},
							Operator: "+",
							Right:    &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "second"}, Value: "second"},
						},
					},
				},
			},
		},
		{
			name: "success - expression statement - function declaration 1",
			input: `
			sun_liyo_tau some_func ne_bana_diye rasoi_mein_bata_diye(x) {
				laadle_ye_le x
			};
			`,
			expectedErrors: []string{},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.LetStatement{
						Token: token.Token{Type: token.LET, Literal: "let"},
						Name:  &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "some_func"}, Value: "some_func"},
						Value: &ast.FunctionLiteral{
							Token: token.Token{Type: token.FUNCTION, Literal: "func"},
							Parameters: []*ast.Identifier{
								{Token: token.Token{Type: token.IDENTIFIER, Literal: "x"}, Value: "x"},
							},
							Body: &ast.BlockStatement{
								Token: token.Token{Type: token.LEFT_BRACE, Literal: "{"},
								Statements: []ast.Statement{
									&ast.ReturnStatement{
										Token:       token.Token{Type: token.RETURN, Literal: "return"},
										ReturnValue: &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "x"}, Value: "x"},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "success - expression statement - function declaration 2",
			input: `
			sun_liyo_tau first ne_bana_diye 10;
			sun_liyo_tau second ne_bana_diye 10;
			sun_liyo_tau third ne_bana_diye 10;
			
			sun_liyo_tau ourFunction ne_bana_diye rasoi_mein_bata_diye(first) {
			  sun_liyo_tau second ne_bana_diye 20;
			
			  first + second + third;
			};
			
			ourFunction(20) + first + second;
			`,
			expectedErrors: []string{},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.LetStatement{
						Token: token.Token{Type: token.LET, Literal: "let"},
						Name:  &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "first"}, Value: "first"},
						Value: &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "10"}, Value: 10},
					},
					&ast.LetStatement{
						Token: token.Token{Type: token.LET, Literal: "let"},
						Name:  &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "second"}, Value: "second"},
						Value: &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "10"}, Value: 10},
					},
					&ast.LetStatement{
						Token: token.Token{Type: token.LET, Literal: "let"},
						Name:  &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "third"}, Value: "third"},
						Value: &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "10"}, Value: 10},
					},
					&ast.LetStatement{
						Token: token.Token{Type: token.LET, Literal: "let"},
						Name:  &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "ourFunction"}, Value: "ourFunction"},
						Value: &ast.FunctionLiteral{
							Token: token.Token{Type: token.FUNCTION, Literal: "func"},
							Parameters: []*ast.Identifier{
								{Token: token.Token{Type: token.IDENTIFIER, Literal: "first"}, Value: "first"},
							},
							Body: &ast.BlockStatement{
								Token: token.Token{Type: token.LEFT_BRACE, Literal: "{"},
								Statements: []ast.Statement{
									&ast.LetStatement{
										Token: token.Token{Type: token.LET, Literal: "let"},
										Name:  &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "second"}, Value: "second"},
										Value: &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "20"}, Value: 20},
									},
									&ast.ExpressionStatement{
										Token: token.Token{Type: token.IDENTIFIER, Literal: "first"},
										Expression: &ast.InfixExpression{
											Token: token.Token{Type: token.ADDITION, Literal: "+"},
											Left: &ast.InfixExpression{
												Token:    token.Token{Type: token.ADDITION, Literal: "+"},
												Left:     &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "first"}, Value: "first"},
												Operator: "+",
												Right:    &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "second"}, Value: "second"},
											},
											Operator: "+",
											Right:    &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "third"}, Value: "third"},
										},
									},
								},
							},
						},
					},
					&ast.ExpressionStatement{
						Token: token.Token{Type: token.IDENTIFIER, Literal: "ourFunction"},
						Expression: &ast.InfixExpression{
							Token: token.Token{Type: token.ADDITION, Literal: "+"},
							Left: &ast.InfixExpression{
								Token: token.Token{Type: token.ADDITION, Literal: "+"},
								Left: &ast.CallExpression{
									Token:    token.Token{Type: token.LEFT_PAREN, Literal: "("},
									Function: &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "ourFunction"}, Value: "ourFunction"},
									Arguments: []ast.Expression{
										&ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "20"}, Value: 20},
									},
								},
								Operator: "+",
								Right:    &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "first"}, Value: "first"},
							},
							Operator: "+",
							Right:    &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "second"}, Value: "second"},
						},
					},
				},
			},
		},
		{
			name:           "success - expression statement - standard expression",
			input:          `5 + 5`,
			expectedErrors: []string{},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: token.Token{Type: token.NUMBER, Literal: "5"},
						Expression: &ast.InfixExpression{
							Token:    token.Token{Type: token.ADDITION, Literal: "+"},
							Left:     &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "5"}, Value: 5},
							Operator: "+",
							Right:    &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "5"}, Value: 5},
						},
					},
				},
			},
		},
		{
			name:           "success - expression statement - standard expression with identifier",
			input:          `some_var + 5`,
			expectedErrors: []string{},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: token.Token{Type: token.IDENTIFIER, Literal: "some_var"},
						Expression: &ast.InfixExpression{
							Token:    token.Token{Type: token.ADDITION, Literal: "+"},
							Left:     &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "some_var"}, Value: "some_var"},
							Operator: "+",
							Right:    &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "5"}, Value: 5},
						},
					},
				},
			},
		},
		{
			name:           "success - expression statement - standard expression with identifier and grouped expression",
			input:          `3 + y * (some_var + 5)`,
			expectedErrors: []string{},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: token.Token{Type: token.NUMBER, Literal: "3"},
						Expression: &ast.InfixExpression{
							Token:    token.Token{Type: token.ADDITION, Literal: "+"},
							Left:     &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "3"}, Value: 3},
							Operator: "+",
							Right: &ast.InfixExpression{
								Token:    token.Token{Type: token.MULTIPLICATION, Literal: "*"},
								Left:     &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "y"}, Value: "y"},
								Operator: "*",
								Right: &ast.InfixExpression{
									Token:    token.Token{Type: token.ADDITION, Literal: "+"},
									Left:     &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "some_var"}, Value: "some_var"},
									Operator: "+",
									Right:    &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "5"}, Value: 5},
								},
							},
						},
					},
				},
			},
		},
		{
			name:           "success - expression statement - standard expression with identifier and prefix operator",
			input:          `-5 + some_var + 5`,
			expectedErrors: []string{},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: token.Token{Type: token.SUBTRACTION, Literal: "-"},
						Expression: &ast.InfixExpression{
							Token: token.Token{Type: token.ADDITION, Literal: "+"},
							Left: &ast.InfixExpression{
								Token: token.Token{Type: token.ADDITION, Literal: "+"},
								Left: &ast.PrefixExpression{
									Token:    token.Token{Type: token.SUBTRACTION, Literal: "-"},
									Operator: "-",
									Operand:  &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "5"}, Value: 5},
								},
								Operator: "+",
								Right:    &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "some_var"}, Value: "some_var"},
							},
							Operator: "+",
							Right:    &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "5"}, Value: 5},
						},
					},
				},
			},
		},
		{
			name:           "success - expression statement - standard expression with identifier and bang prefix operator",
			input:          `!5 + some_var + 5`,
			expectedErrors: []string{},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: token.Token{Type: token.BANG, Literal: "!"},
						Expression: &ast.InfixExpression{
							Token: token.Token{Type: token.ADDITION, Literal: "+"},
							Left: &ast.InfixExpression{
								Token: token.Token{Type: token.ADDITION, Literal: "+"},
								Left: &ast.PrefixExpression{
									Token:    token.Token{Type: token.BANG, Literal: "!"},
									Operator: "!",
									Operand:  &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "5"}, Value: 5},
								},
								Operator: "+",
								Right:    &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "some_var"}, Value: "some_var"},
							},
							Operator: "+",
							Right:    &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "5"}, Value: 5},
						},
					},
				},
			},
		},
		{
			name: "success - if / else statement",
			input: `
			agar_maan_lo (x <= 3) {
				laadle_ye_le a;
			} na_toh {
				laadle_ye_le b;
			}
			`,
			expectedErrors: []string{},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: token.Token{Type: token.IF, Literal: "if"},
						Expression: &ast.ConditionalExpression{
							Token: token.Token{Type: token.IF, Literal: "if"},
							Condition: &ast.InfixExpression{
								Token:    token.Token{Type: token.LESSER_EQUALS, Literal: "<="},
								Left:     &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "x"}, Value: "x"},
								Operator: "<=",
								Right:    &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "3"}, Value: 3},
							},
							Consequence: &ast.BlockStatement{
								Token: token.Token{Type: token.LEFT_BRACE, Literal: "{"},
								Statements: []ast.Statement{
									&ast.ReturnStatement{
										Token:       token.Token{Type: token.RETURN, Literal: "return"},
										ReturnValue: &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "a"}, Value: "a"},
									},
								},
							},
							Alternative: &ast.BlockStatement{
								Token: token.Token{Type: token.LEFT_BRACE, Literal: "{"},
								Statements: []ast.Statement{
									&ast.ReturnStatement{
										Token:       token.Token{Type: token.RETURN, Literal: "return"},
										ReturnValue: &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "b"}, Value: "b"},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "success - if statement without else block followed by another statement",
			input: `
			agar_maan_lo (x <= 3) {
				laadle_ye_le a;
			}
			sun_liyo_tau x ne_bana_diye 5;
			`,
			expectedErrors: []string{},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: token.Token{Type: token.IF, Literal: "if"},
						Expression: &ast.ConditionalExpression{
							Token: token.Token{Type: token.IF, Literal: "if"},
							Condition: &ast.InfixExpression{
								Token:    token.Token{Type: token.LESSER_EQUALS, Literal: "<="},
								Left:     &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "x"}, Value: "x"},
								Operator: "<=",
								Right:    &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "3"}, Value: 3},
							},
							Consequence: &ast.BlockStatement{
								Token: token.Token{Type: token.LEFT_BRACE, Literal: "{"},
								Statements: []ast.Statement{
									&ast.ReturnStatement{
										Token:       token.Token{Type: token.RETURN, Literal: "return"},
										ReturnValue: &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "a"}, Value: "a"},
									},
								},
							},
							Alternative: nil,
						},
					},
					&ast.LetStatement{
						Token: token.Token{Type: token.LET, Literal: "let"},
						Name:  &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "x"}, Value: "x"},
						Value: &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "5"}, Value: 5},
					},
				},
			},
		},
		{
			name: "success - while loop with break and continue",
			input: `
			jab_tak (x <= 3) {
				agar_maan_lo (x == 3) {
					rok_diye;
				} na_toh {
					jaan_de;
				}
			}
			`,
			expectedErrors: []string{},
			expectedProgram: &ast.Program{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: token.Token{Type: token.WHILE, Literal: "while"},
						Expression: &ast.WhileLoopExpression{
							Token: token.Token{Type: token.WHILE, Literal: "while"},
							Condition: &ast.InfixExpression{
								Token:    token.Token{Type: token.LESSER_EQUALS, Literal: "<="},
								Left:     &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "x"}, Value: "x"},
								Operator: "<=",
								Right:    &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "3"}, Value: 3},
							},
							Body: &ast.BlockStatement{
								Token: token.Token{Type: token.LEFT_BRACE, Literal: "{"},
								Statements: []ast.Statement{
									&ast.ExpressionStatement{
										Token: token.Token{Type: token.IF, Literal: "if"},
										Expression: &ast.ConditionalExpression{
											Token: token.Token{Type: token.IF, Literal: "if"},
											Condition: &ast.InfixExpression{
												Token:    token.Token{Type: token.EQUALS, Literal: "=="},
												Left:     &ast.Identifier{Token: token.Token{Type: token.IDENTIFIER, Literal: "x"}, Value: "x"},
												Operator: "==",
												Right:    &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Literal: "3"}, Value: 3},
											},
											Consequence: &ast.BlockStatement{
												Token: token.Token{Type: token.LEFT_BRACE, Literal: "{"},
												Statements: []ast.Statement{
													&ast.ExpressionStatement{
														Token:      token.Token{Type: token.BREAK, Literal: "break"},
														Expression: &ast.BreakExpression{Token: token.Token{Type: token.BREAK, Literal: "break"}},
													},
												},
											},
											Alternative: &ast.BlockStatement{
												Token: token.Token{Type: token.LEFT_BRACE, Literal: "{"},
												Statements: []ast.Statement{
													&ast.ExpressionStatement{
														Token:      token.Token{Type: token.CONTINUE, Literal: "continue"},
														Expression: &ast.ContinueExpression{Token: token.Token{Type: token.CONTINUE, Literal: "continue"}},
													},
												},
											},
										},
									},
								},
							},
						},
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
