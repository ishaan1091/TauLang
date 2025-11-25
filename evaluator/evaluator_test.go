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
			input: "tau_ka_jugaad(x) { x + 2; };",
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
			input:          "sun_liyo_tau add ne_bana_diye tau_ka_jugaad(x, y) { x + y; }; add(5 + 5, add(5, 5));",
			expectedObject: &object.Integer{Value: 20},
		},
		{
			name:           "success - call expression 2",
			input:          "tau_ka_jugaad(x) { x; }(5)",
			expectedObject: &object.Integer{Value: 5},
		},
		{
			name: "success - call expression 3",
			input: `
			sun_liyo_tau first ne_bana_diye 10;
			sun_liyo_tau second ne_bana_diye 10;
			sun_liyo_tau third ne_bana_diye 10;
			
			sun_liyo_tau ourFunction ne_bana_diye tau_ka_jugaad(first) {
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
			sun_liyo_tau fn ne_bana_diye tau_ka_jugaad() { 
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
			sun_liyo_tau fn ne_bana_diye tau_ka_jugaad() { 
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
			sun_liyo_tau power ne_bana_diye tau_ka_jugaad(x, n) { 
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
		{
			name: "success - while loop expression - break statement 1",
			input: `
			sun_liyo_tau incr ne_bana_diye tau_ka_jugaad(n) { 
				sun_liyo_tau i ne_bana_diye 0;
				jab_tak (i) {
					i ne_bana_diye i + 1
					
					agar_maan_lo (i == n) {
						rok_diye
					}
				}
				laadle_ye_le i; 
			}; 
			incr(5) + incr(100)
			`,
			expectedObject: &object.Integer{Value: 105},
		},
		{
			name: "success - while loop expression - break statement 2",
			input: `
			sun_liyo_tau power ne_bana_diye tau_ka_jugaad(x, n) { 
				sun_liyo_tau i ne_bana_diye 0;
				sun_liyo_tau result ne_bana_diye 1;
				jab_tak (i) {
					result ne_bana_diye x * result;
					i ne_bana_diye i + 1
					
					agar_maan_lo (i == n) {
						rok_diye
					}
				}
				laadle_ye_le result; 
			}; 
			power(2, 5) + power(10, 4)
			`,
			expectedObject: &object.Integer{Value: 10032},
		},
		{
			name: "failure - break statement outside while loop",
			input: `
			agar_maan_lo (saccha) {
				rok_diye
			}
			`,
			expectedObject: &object.Error{Message: "found break statement outside of loop"},
		},
		{
			name: "success - while loop expression - continue statement",
			input: `
			sun_liyo_tau rangeSum ne_bana_diye tau_ka_jugaad(l, r) { 
				sun_liyo_tau i ne_bana_diye 0;
				sun_liyo_tau result ne_bana_diye 0;
				jab_tak (i <= r) {
					agar_maan_lo (i < l) {
						i ne_bana_diye i + 1;
						jaan_de;
					}
					result ne_bana_diye result + i;
					i ne_bana_diye i + 1
				}
				laadle_ye_le result; 
			}; 
			rangeSum(5, 10)
			`,
			expectedObject: &object.Integer{Value: 45},
		},
		{
			name: "failure - continue statement outside while loop",
			input: `
			agar_maan_lo (saccha) {
				jaan_de
			}
			`,
			expectedObject: &object.Error{Message: "found continue statement outside of loop"},
		},
		{
			name:           "success - builtin function - len - string 1",
			input:          "sun_liyo_tau a ne_bana_diye \"test string\"; len(a);",
			expectedObject: &object.Integer{Value: 11},
		},
		{
			name:           "success - builtin function - len - string 2",
			input:          "len(\"twitter\");",
			expectedObject: &object.Integer{Value: 7},
		},
		{
			name:           "success - builtin function - len - array",
			input:          "len([1, \"twitter\", saccha]);",
			expectedObject: &object.Integer{Value: 3},
		},
		{
			name:           "success - builtin function - len - hashmap",
			input:          "len({\"one\":1, 2: 2, saccha: 3});",
			expectedObject: &object.Integer{Value: 3},
		},
		{
			name:  "success - array literal",
			input: "[3, \"hello\", saccha, tau_ka_jugaad(x) { x + 2; }]",
			expectedObject: &object.Array{
				Elements: []object.Object{
					&object.Integer{Value: 3},
					&object.String{Value: "hello"},
					&object.Boolean{Value: true},
					&object.Function{
						Params: []*ast.Identifier{
							{Token: token.Token{Type: token.IDENTIFIER, Literal: "x"}, Value: "x"},
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
			},
		},
		{
			name:           "success - index expression - array 1",
			input:          "[1, 2, 3][2]",
			expectedObject: &object.Integer{Value: 3},
		},
		{
			name:           "success - index expression - array 2",
			input:          "sun_liyo_tau i ne_bana_diye 0; [1][i];",
			expectedObject: &object.Integer{Value: 1},
		},
		{
			name:           "success - index expression - array 3",
			input:          "sun_liyo_tau myArray ne_bana_diye [1, 2, 3]; sun_liyo_tau i ne_bana_diye myArray[0]; myArray[i]",
			expectedObject: &object.Integer{Value: 2},
		},
		{
			name:           "success - index expression - array 4",
			input:          "[1, 2, 3][1 + 1];",
			expectedObject: &object.Integer{Value: 3},
		},
		{
			name:           "success - index expression - array 5",
			input:          "[1, 2, 3][3]",
			expectedObject: &object.Null{},
		},
		{
			name:           "success - index expression - array 6",
			input:          "[1, 2, 3][-1]",
			expectedObject: &object.Null{},
		},
		{
			name:           "success - index expression - hashmap 1",
			input:          `{"foo": 5}["foo"]`,
			expectedObject: &object.Integer{Value: 5},
		},
		{
			name:           "success - index expression - hashmap 2",
			input:          `{"foo": 5}["bar"]`,
			expectedObject: &object.Null{},
		},
		{
			name:           "success - index expression - hashmap 3",
			input:          `sun_liyo_tau key ne_bana_diye "foo"; {"foo": 5}[key]`,
			expectedObject: &object.Integer{Value: 5},
		},
		{
			name:           "success - index expression - hashmap 4",
			input:          `{}["foo"]`,
			expectedObject: &object.Null{},
		},
		{
			name:           "success - index expression - hashmap 5",
			input:          `{5: 5}[5]`,
			expectedObject: &object.Integer{Value: 5},
		},
		{
			name:           "success - index expression - hashmap 6",
			input:          `{saccha: 5}[saccha]`,
			expectedObject: &object.Integer{Value: 5},
		},
		{
			name:           "success - index expression - hashmap 7",
			input:          `{jhootha: 5}[jhootha]`,
			expectedObject: &object.Integer{Value: 5},
		},
		{
			name:           "failure - index expression - array",
			input:          "[1, 2, 3][saccha]",
			expectedObject: &object.Error{Message: "index operator not supported: ARRAY[BOOLEAN]"},
		},
		{
			name:           "failure - index expression - hashmap",
			input:          "{1: 5}[tau_ka_jugaad(x) {x}]",
			expectedObject: &object.Error{Message: "unusable as hash key: FUNCTION"},
		},
		{
			name:           "success - builtin function - first",
			input:          "first([1, 2, 3])",
			expectedObject: &object.Integer{Value: 1},
		},
		{
			name:           "success - builtin function - last",
			input:          "last([1, 2, 3])",
			expectedObject: &object.Integer{Value: 3},
		},
		{
			name:           "success - builtin function - last",
			input:          "last(push([1, 2, 3], 4))",
			expectedObject: &object.Integer{Value: 4},
		},
		{
			name: "success - hashmap",
			input: `sun_liyo_tau two ne_bana_diye "two";
			{
				"one": 10 - 9,
				two: 1 + 1,
				"thr" + "ee": 6 / 2,
				4: 4,
				saccha: 5,
				jhootha: 6
			}`,
			expectedObject: &object.HashMap{
				Pairs: map[object.HashKey]object.HashPair{
					(&object.String{Value: "one"}).Hash(): {
						Key:   &object.String{Value: "one"},
						Value: &object.Integer{Value: 1},
					},
					(&object.String{Value: "two"}).Hash(): {
						Key:   &object.String{Value: "two"},
						Value: &object.Integer{Value: 2},
					},
					(&object.String{Value: "three"}).Hash(): {
						Key:   &object.String{Value: "three"},
						Value: &object.Integer{Value: 3},
					},
					(&object.Integer{Value: 4}).Hash(): {
						Key:   &object.Integer{Value: 4},
						Value: &object.Integer{Value: 4},
					},
					evaluator.TRUE.Hash(): {
						Key:   &object.Boolean{Value: true},
						Value: &object.Integer{Value: 5},
					},
					evaluator.FALSE.Hash(): {
						Key:   &object.Boolean{Value: false},
						Value: &object.Integer{Value: 6},
					},
				},
			},
		},
		{
			name:           "failure - hashmap",
			input:          `{tau_ka_jugaad(x) { x; }: 5}`,
			expectedObject: &object.Error{Message: "unusable as hash key: FUNCTION"},
		},
		{
			name: "success - hash index assignment",
			input: `sun_liyo_tau map ne_bana_diye {"one": 1, "two": 2};
			map["three"] ne_bana_diye 3;
			map["three"];`,
			expectedObject: &object.Integer{Value: 3},
		},
		{
			name: "success - hash index assignment - overwrite",
			input: `sun_liyo_tau map ne_bana_diye {"one": 1, "two": 2};
			map["one"] ne_bana_diye 10;
			map["one"];`,
			expectedObject: &object.Integer{Value: 10},
		},
		{
			name: "success - array index assignment",
			input: `sun_liyo_tau arr ne_bana_diye [1, 2, 3];
			arr[1] ne_bana_diye 20;
			arr[1];`,
			expectedObject: &object.Integer{Value: 20},
		},
		{
			name: "success - array index assignment - extend",
			input: `sun_liyo_tau arr ne_bana_diye [1, 2];
			arr[5] ne_bana_diye 10;
			arr[5];`,
			expectedObject: &object.Integer{Value: 10},
		},
		{
			name: "success - array index assignment - verify other elements",
			input: `sun_liyo_tau arr ne_bana_diye [1, 2];
			arr[5] ne_bana_diye 10;
			arr[0];`,
			expectedObject: &object.Integer{Value: 1},
		},
		{
			name:           "failure - index assignment on undefined variable",
			input:          `map["key"] ne_bana_diye 1;`,
			expectedObject: &object.Error{Message: "identifier not found: map"},
		},
		{
			name:           "failure - index assignment on non-indexable type",
			input:          `sun_liyo_tau x ne_bana_diye 5; x[0] ne_bana_diye 1;`,
			expectedObject: &object.Error{Message: "index assignment not supported for type: INTEGER"},
		},
		{
			name:           "success - no executable code only comments",
			input:          `// sun_liyo_tau x ne_bana_diye 5; x[0] ne_bana_diye 1;`,
			expectedObject: &object.Null{},
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
