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
			name:  "success - integer",
			input: "5;",
			expectedObject: &object.Integer{
				Value: 5,
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
			assert.NotNil(t, program)

			o := evaluator.Eval(program)
			assert.Equal(t, tc.expectedObject, o)
		})
	}
}
