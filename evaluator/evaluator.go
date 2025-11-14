package evaluator

import (
	"fmt"
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
		return evalProgram(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return getBoolObject(node.Value)
	case *ast.PrefixExpression:
		return evalPrefixExpression(node.Operator, node.Operand)
	}
	return nil
}

func evalProgram(statements []ast.Statement) object.Object {
	var result object.Object
	for _, s := range statements {
		result = Eval(s)

		if result.Type() == object.ERROR_OBJ {
			return result
		}
	}
	return result
}

func getBoolObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func evalPrefixExpression(operator string, operand ast.Expression) object.Object {
	evaluatedOperand := Eval(operand)
	switch operator {
	case "-":
		return evalMinusPrefixOperatorExpression(evaluatedOperand)
	case "!":
		return evalBangOperatorExpression(evaluatedOperand)
	default:
		return newError("unknown prefix expression: %s%s", operator, evaluatedOperand.Type())
	}
}

func evalBangOperatorExpression(operand object.Object) object.Object {
	if operand == FALSE {
		return TRUE
	}
	return FALSE
}

func evalMinusPrefixOperatorExpression(operand object.Object) object.Object {
	if operand.Type() != object.INTEGER_OBJ {
		return newError("unknown operator: -%s", operand.Type())
	}

	value := operand.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func newError(messageTemplate string, args ...any) *object.Error {
	return &object.Error{Message: fmt.Sprintf(messageTemplate, args...)}
}
