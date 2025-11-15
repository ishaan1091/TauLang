package evaluator_test

import (
	"taulang/evaluator"
	"taulang/lexer"
	"taulang/object"
	"taulang/parser"
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

			o := evaluator.Eval(program)
			assert.Equal(t, tc.expectedObject, o)
		})
	}
}
