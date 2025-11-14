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
