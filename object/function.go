package object

import (
	"strings"
	"taulang/ast"
)

type Function struct {
	Params []*ast.Identifier
	Body   *ast.BlockStatement
	Env    Environment
}

func (f *Function) Type() Type {
	return FUNCTION_OBJ
}

func (f *Function) Inspect() string {
	var out strings.Builder
	var params []string
	for _, p := range f.Params {
		params = append(params, p.String())
	}
	out.WriteString("func")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}
