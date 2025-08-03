package lexer

import (
	"jsonparser/cmd/jsonparser/commons"
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name      string
		content   string
		expectErr bool
		err       error
		result    []*commons.Token
	}{
		{
			name:      "success returns nil result",
			content:   "",
			expectErr: false,
			err:       nil,
			result:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Tokenize(tt.content)

			if tt.expectErr {
				if err != tt.err {
					t.Error("error not matching expected error")
				}
			} else {
				if err != nil {
					t.Error("error not expected")
				}

				if !reflect.DeepEqual(result, tt.result) {
					t.Error("result not matching expected result")
				}
			}
		})
	}
}
