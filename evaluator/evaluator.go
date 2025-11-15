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
	case *ast.InfixExpression:
		return evalInfixExpression(node.Operator, node.Left, node.Right)
	default:
		return newError("no defined evaluations for input: %s", node.String())
	}
}

func evalProgram(statements []ast.Statement) object.Object {
	var result object.Object
	for _, s := range statements {
		result = Eval(s)

		if isError(result) {
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
	if isError(evaluatedOperand) {
		return evaluatedOperand
	}

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

func evalInfixExpression(operator string, left ast.Expression, right ast.Expression) object.Object {
	evaluatedLeft := Eval(left)
	if isError(evaluatedLeft) {
		return evaluatedLeft
	}

	evaluatedRight := Eval(right)
	if isError(evaluatedRight) {
		return evaluatedRight
	}

	switch {
	case evaluatedLeft.Type() == object.INTEGER_OBJ && evaluatedRight.Type() == object.INTEGER_OBJ:
		return evaluateIntegerInfixExpression(operator, evaluatedLeft.(*object.Integer), evaluatedRight.(*object.Integer))

	// Below we are doing direct object comparisons of evaluated results as at this point we expect
	// only boolean types which are using same object for all TRUE / FALSE
	case operator == "==":
		return getBoolObject(evaluatedLeft == evaluatedRight)
	case operator == "!=":
		return getBoolObject(evaluatedLeft != evaluatedRight)

	case evaluatedLeft.Type() != evaluatedRight.Type():
		return newError("type mismatch: %s %s %s", evaluatedLeft.Type(), operator, evaluatedRight.Type())
	default:
		return newError("unknown operator: %s %s %s", evaluatedLeft.Type(), operator, evaluatedRight.Type())
	}
}

func evaluateIntegerInfixExpression(operator string, left *object.Integer, right *object.Integer) object.Object {
	leftVal := left.Value
	rightVal := right.Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		if rightVal == 0 {
			return newError("division by zero")
		}

		return &object.Integer{Value: leftVal / rightVal}
	case "==":
		return &object.Boolean{Value: leftVal == rightVal}
	case "!=":
		return &object.Boolean{Value: leftVal != rightVal}
	case "<":
		return &object.Boolean{Value: leftVal < rightVal}
	case "<=":
		return &object.Boolean{Value: leftVal <= rightVal}
	case ">":
		return &object.Boolean{Value: leftVal > rightVal}
	case ">=":
		return &object.Boolean{Value: leftVal >= rightVal}
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func newError(messageTemplate string, args ...any) *object.Error {
	return &object.Error{Message: fmt.Sprintf(messageTemplate, args...)}
}

func isError(obj object.Object) bool {
	return obj != nil && obj.Type() == object.ERROR_OBJ
}
