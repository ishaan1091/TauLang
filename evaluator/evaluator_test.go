package evaluator_test

import (
	"taulang/ast"
	"taulang/evaluator"
	"taulang/lexer"
	"taulang/object"
	"taulang/parser"
	"taulang/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluator(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedObject object.Object
	}{
		{
			name:  "success - integer 1",
			input: "5;",
			expectedObject: &object.Integer{
				Value: 5,
			},
		},
		{
			name:  "success - integer 2",
			input: "10;",
			expectedObject: &object.Integer{
				Value: 10,
			},
		},
		{
			name:  "success - boolean true",
			input: "saccha;",
			expectedObject: &object.Boolean{
				Value: true,
			},
		},
		{
			name:  "success - boolean false",
			input: "jhootha;",
			expectedObject: &object.Boolean{
				Value: false,
			},
		},
		{
			name:  "success - prefix expression - bang true",
			input: "!saccha;",
			expectedObject: &object.Boolean{
				Value: false,
			},
		},
		{
			name:  "success - boolean expression - bang false",
			input: "!jhootha;",
			expectedObject: &object.Boolean{
				Value: true,
			},
		},
		{
			name:  "success - boolean expression - bang integer",
			input: "!!!!-5;",
			expectedObject: &object.Boolean{
				Value: true,
			},
		},
		{
			name:  "success - prefix expression - minus 1",
			input: "-5;",
			expectedObject: &object.Integer{
				Value: -5,
			},
		},
		{
			name:  "success - prefix expression - minus 2",
			input: "-10;",
			expectedObject: &object.Integer{
				Value: -10,
			},
		},
		{
			name:           "failure - prefix expression - minus operator on non integer types 1",
			input:          "-saccha;",
			expectedObject: &object.Error{Message: "unknown operator: -BOOLEAN"},
		},
		{
			name:           "failure - prefix expression - minus operator on non integer types 2",
			input:          "-jhootha;",
			expectedObject: &object.Error{Message: "unknown operator: -BOOLEAN"},
		},
		{
			name:           "failure - multiple statements with error in first",
			input:          "-jhootha;5;",
			expectedObject: &object.Error{Message: "unknown operator: -BOOLEAN"},
		},
		{
			name:           "success - infix expression 1",
			input:          "5 + 5 + 5 + 5 - 10;",
			expectedObject: &object.Integer{Value: 10},
		},
		{
			name:           "success - infix expression 2",
			input:          "2 * 2 * 2 * 2 * 2;",
			expectedObject: &object.Integer{Value: 32},
		},
		{
			name:           "success - infix expression 3",
			input:          "-50 + 100 + -50;",
			expectedObject: &object.Integer{Value: 0},
		},
		{
			name:           "success - infix expression 4",
			input:          "5 * 2 + 10;",
			expectedObject: &object.Integer{Value: 20},
		},
		{
			name:           "success - infix expression 5",
			input:          "5 + 2 * 10",
			expectedObject: &object.Integer{Value: 25},
		},
		{
			name:           "success - infix expression 6",
			input:          "20 + 2 * -10",
			expectedObject: &object.Integer{Value: 0},
		},
		{
			name:           "success - infix expression 7",
			input:          "50 / 2 * 2 + 10;",
			expectedObject: &object.Integer{Value: 60},
		},
		{
			name:           "success - infix expression 8",
			input:          "2 * (5 + 10);",
			expectedObject: &object.Integer{Value: 30},
		},
		{
			name:           "success - infix expression 9",
			input:          "3 * 3 * 3 + 10",
			expectedObject: &object.Integer{Value: 37},
		},
		{
			name:           "success - infix expression 10",
			input:          "3 * (3 * 3) + 10",
			expectedObject: &object.Integer{Value: 37},
		},
		{
			name:           "success - infix expression 11",
			input:          "(5 + 10 * 2 + 15 / 3) * 2 + -10",
			expectedObject: &object.Integer{Value: 50},
		},
		{
			name:           "failure - infix expression - division by 0",
			input:          "2 / 0 + -10",
			expectedObject: &object.Error{Message: "division by zero"},
		},
		{
			name:           "failure - infix expression type mismatch",
			input:          "5 + saccha; 5;",
			expectedObject: &object.Error{Message: "type mismatch: INTEGER + BOOLEAN"},
		},
		{
			name:           "failure - infix expression type mismatch",
			input:          "5; saccha + jhootha; 5",
			expectedObject: &object.Error{Message: "unknown operator: BOOLEAN + BOOLEAN"},
		},
		{
			name:           "success - infix expression - truthy equality object comparison",
			input:          "!!2 == saccha",
			expectedObject: &object.Boolean{Value: true},
		},
		{
			name:           "success - infix expression - falsy equality object comparison",
			input:          "2 == jhootha",
			expectedObject: &object.Boolean{Value: false},
		},
		{
			name:           "success - infix expression - falsy inequality object comparison",
			input:          "!!2 != saccha",
			expectedObject: &object.Boolean{Value: false},
		},
		{
			name:           "success - infix expression - truthy inequality object comparison",
			input:          "2 != jhootha",
			expectedObject: &object.Boolean{Value: true},
		},
		{
			name:           "success - infix expression - integer comparison 1",
			input:          "2 == 10 - 8",
			expectedObject: &object.Boolean{Value: true},
		},
		{
			name:           "success - infix expression - integer comparison 2",
			input:          "2 == 10 * 8",
			expectedObject: &object.Boolean{Value: false},
		},
		{
			name:           "success - infix expression - integer comparison 3",
			input:          "2 != 10 - 8",
			expectedObject: &object.Boolean{Value: false},
		},
		{
			name:           "success - infix expression - integer comparison 4",
			input:          "2 != 10",
			expectedObject: &object.Boolean{Value: true},
		},
		{
			name:           "success - infix expression - integer comparison 5",
			input:          "2 < 10",
			expectedObject: &object.Boolean{Value: true},
		},
		{
			name:           "success - infix expression - integer comparison 6",
			input:          "2 < 1",
			expectedObject: &object.Boolean{Value: false},
		},
		{
			name:           "success - infix expression - integer comparison 7",
			input:          "2 <= 2",
			expectedObject: &object.Boolean{Value: true},
		},
		{
			name:           "success - infix expression - integer comparison 8",
			input:          "2 <= 1",
			expectedObject: &object.Boolean{Value: false},
		},
		{
			name:           "success - infix expression - integer comparison 9",
			input:          "2 > 10 - 8",
			expectedObject: &object.Boolean{Value: false},
		},
		{
			name:           "success - infix expression - integer comparison 10",
			input:          "2 > 1",
			expectedObject: &object.Boolean{Value: true},
		},
		{
			name:           "success - infix expression - integer comparison 11",
			input:          "2 >= 10 - 8",
			expectedObject: &object.Boolean{Value: true},
		},
		{
			name:           "success - infix expression - integer comparison 12",
			input:          "2 >= 10",
			expectedObject: &object.Boolean{Value: false},
		},
		{
			name:           "success - conditional expression 1",
			input:          "agar_maan_lo (1 < 2) { 10 }",
			expectedObject: &object.Integer{Value: 10},
		},
		{
			name:           "success - conditional expression 2",
			input:          "agar_maan_lo (1 > 2) { 10 }",
			expectedObject: &object.Null{},
		},
		{
			name:           "success - conditional expression 3",
			input:          "agar_maan_lo (1 > 2) { 10 } na_toh { 20 }",
			expectedObject: &object.Integer{Value: 20},
		},
		{
			name:           "success - conditional expression 4",
			input:          "agar_maan_lo (1 < 2) { 10 } na_toh { 20 }",
			expectedObject: &object.Integer{Value: 10},
		},
		{
			name:           "success - return statement 1",
			input:          "laadle_ye_le 2 * 5; 9;",
			expectedObject: &object.Integer{Value: 10},
		},
		{
			name:           "success - return statement 2",
			input:          "9; laadle_ye_le 2 * 3; 9;",
			expectedObject: &object.Integer{Value: 6},
		},
		{
			name:           "success - return statement 3",
			input:          "agar_maan_lo (10 > 1) { laadle_ye_le 10; }",
			expectedObject: &object.Integer{Value: 10},
		},
		{
			name: "success - return statement 4",
			input: `
			agar_maan_lo (10 > 1) {
			  agar_maan_lo (10 > 1) {
				laadle_ye_le 10;
			  }
			
			  laadle_ye_le 1;
			}
			`,
			expectedObject: &object.Integer{Value: 10},
		},
		{
			name:           "success - string literal expression",
			input:          "\"test_string\";",
			expectedObject: &object.String{Value: "test_string"},
		},
		{
			name:           "success - infix expression - string concatenation",
			input:          "laadle_ye_le \"tau\" + \" khush\";",
			expectedObject: &object.String{Value: "tau khush"},
		},
		{
			name:           "success - infix expression - truthy string equality check",
			input:          "\"tau\" == \"tau\"",
			expectedObject: &object.Boolean{Value: true},
		},
		{
			name:           "success - infix expression - falsy string equality check",
			input:          "\"tau\" == \"not_tau\"",
			expectedObject: &object.Boolean{Value: false},
		},
		{
			name:           "failure - infix expression - string unknown operator",
			input:          "\"tau\" * \"not_tau\";",
			expectedObject: &object.Error{Message: "unknown operator: STRING * STRING"},
		},
		{
			name:           "success - let statement 1",
			input:          "sun_liyo_tau a ne_bana_diye 5; a;",
			expectedObject: &object.Integer{Value: 5},
		},
		{
			name: "success - let statement 2",
			input: `
			sun_liyo_tau a ne_bana_diye 5 * 2
			sun_liyo_tau b ne_bana_diye a + 2;
			sun_liyo_tau c ne_bana_diye a + b + 8
			c;
			`,
			expectedObject: &object.Integer{Value: 30},
		},
		{
			name:           "failure - identifier not found",
			input:          "sun_liyo_tau a ne_bana_diye 5; c;",
			expectedObject: &object.Error{Message: "identifier not found: c"},
		},
		{
			name:           "failure - identifier not found",
			input:          "a; sun_liyo_tau a ne_bana_diye 5;",
			expectedObject: &object.Error{Message: "identifier not found: a"},
		},
		{
			name:  "success - function definition",
			input: "rasoi_mein_bata_diye(x) { x + 2; };",
			expectedObject: &object.Function{
				Params: []*ast.Identifier{
					{
						Token: token.Token{Type: token.IDENTIFIER, Literal: "x"},
						Value: "x",
					},
				},
				Body: &ast.BlockStatement{
					Token: token.Token{Type: token.LEFT_BRACE, Literal: "{"},
					Statements: []ast.Statement{
						&ast.ExpressionStatement{
							Token: token.Token{Type: token.IDENTIFIER, Literal: "x"},
							Expression: &ast.InfixExpression{
								Token: token.Token{Type: token.ADDITION, Literal: "+"},
								Left: &ast.Identifier{
									Token: token.Token{Type: token.IDENTIFIER, Literal: "x"},
									Value: "x",
								},
								Operator: "+",
								Right: &ast.IntegerLiteral{
									Token: token.Token{Type: token.NUMBER, Literal: "2"},
									Value: 2,
								},
							},
						},
					},
				},
				Env: object.NewEnvironment(),
			},
		},
		{
			name:           "success - call expression 1",
			input:          "sun_liyo_tau add ne_bana_diye rasoi_mein_bata_diye(x, y) { x + y; }; add(5 + 5, add(5, 5));",
			expectedObject: &object.Integer{Value: 20},
		},
		{
			name:           "success - call expression 2",
			input:          "rasoi_mein_bata_diye(x) { x; }(5)",
			expectedObject: &object.Integer{Value: 5},
		},
		{
			name: "success - call expression 3",
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
			expectedObject: &object.Integer{Value: 70},
		},
		{
			name:           "failure - call expression - not a function",
			input:          "sun_liyo_tau x ne_bana_diye 65; x(5)",
			expectedObject: &object.Error{Message: "not a function: INTEGER"},
		},
		{
			name:           "success - assignment statement 1",
			input:          "sun_liyo_tau x ne_bana_diye 65; x ne_bana_diye x + 1;x",
			expectedObject: &object.Integer{Value: 66},
		},
		{
			name: "success - assignment statement 2",
			input: `
			sun_liyo_tau fn ne_bana_diye rasoi_mein_bata_diye() { 
				laadle_ye_le 12; 
			}; 
			sun_liyo_tau x ne_bana_diye 65;
			x ne_bana_diye fn();
			x ne_bana_diye x + 1;
			laadle_ye_le x;
			`,
			expectedObject: &object.Integer{Value: 13},
		},
		{
			name: "success - while loop expression 1",
			input: `
			sun_liyo_tau fn ne_bana_diye rasoi_mein_bata_diye() { 
				laadle_ye_le 12; 
			}; 
			sun_liyo_tau x ne_bana_diye 65;
			x ne_bana_diye fn();
			jab_tak (x < 65) {
				x ne_bana_diye x + fn();
			}
			x ne_bana_diye x + 1;
			laadle_ye_le x;
			`,
			expectedObject: &object.Integer{Value: 73},
		},
		{
			name: "success - while loop expression 2",
			input: `
			sun_liyo_tau power ne_bana_diye rasoi_mein_bata_diye(x, n) { 
				sun_liyo_tau i ne_bana_diye 0;
				sun_liyo_tau result ne_bana_diye 1;
				jab_tak (i < n) {
					result ne_bana_diye x * result;
					i ne_bana_diye i + 1
				}
				laadle_ye_le result; 
			}; 
			power(2, 5) + power(10, 4)
			`,
			expectedObject: &object.Integer{Value: 10032},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l, err := lexer.NewLexer(tc.input)
			assert.NoError(t, err)

			p := parser.NewParser(l)
			program := p.Parse()
			assert.NotNil(t, program)
			errors := p.Errors()
			assert.Empty(t, errors)

			env := object.NewEnvironment()
			o := evaluator.Eval(program, env)
			assert.Equal(t, tc.expectedObject, o)
		})
	}
}
