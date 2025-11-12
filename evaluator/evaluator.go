package evaluator

import (
	"taulang/ast"
	"taulang/object"
)

var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		var result object.Object
		for _, s := range node.Statements {
			result = Eval(s)
		}
		return result
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return getBoolObject(node.Value)
	}
	return nil
}

func getBoolObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}
