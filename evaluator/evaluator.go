package evaluator

import (
	"fmt"
	"taulang/ast"
	"taulang/object"
)

var (
	NULL  = &object.Null{}
	BREAK = &object.Break{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node, env object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node.Statements, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return getBoolObject(node.Value)
	case *ast.String:
		return &object.String{Value: node.Value}
	case *ast.PrefixExpression:
		return evalPrefixExpression(node.Operator, node.Operand, env)
	case *ast.InfixExpression:
		return evalInfixExpression(node.Operator, node.Left, node.Right, env)
	case *ast.ConditionalExpression:
		return evalConditionalExpression(node.Condition, node.Consequence, node.Alternative, env)
	case *ast.BlockStatement:
		return evalBlock(node.Statements, env)
	case *ast.ReturnStatement:
		return evalReturnStatement(node.ReturnValue, env)
	case *ast.LetStatement:
		return evalLetStatement(node.Name, node.Value, env)
	case *ast.Identifier:
		return evalIdentifier(node.Value, env)
	case *ast.FunctionLiteral:
		return &object.Function{Params: node.Parameters, Body: node.Body, Env: env}
	case *ast.CallExpression:
		return evaluateCallExpression(node.Function, node.Arguments, env)
	case *ast.AssignmentStatement:
		return evalAssignmentStatement(node.Name, node.Value, env)
	case *ast.WhileLoopExpression:
		return evalWhileLoopExpression(node.Condition, node.Body, env)
	case *ast.BreakStatement:
		return BREAK
	default:
		return newError("no defined evaluations for input: %s", node.String())
	}
}

func evalIdentifier(identifierName string, env object.Environment) object.Object {
	if obj, ok := env.Get(identifierName); ok {
		return obj
	}
	return newError("identifier not found: %s", identifierName)
}

func evalReturnStatement(returnValue ast.Expression, env object.Environment) object.Object {
	evaluatedReturnValue := Eval(returnValue, env)
	if isError(evaluatedReturnValue) {
		return evaluatedReturnValue
	}
	return &object.ReturnValue{Value: evaluatedReturnValue}
}

func evalProgram(statements []ast.Statement, env object.Environment) object.Object {
	var result object.Object
	for _, s := range statements {
		result = Eval(s, env)

		if isError(result) {
			return result
		}

		if isReturnValue(result) {
			return unwrapReturnValue(result)
		}

		if isBreakStatement(result) {
			return newError("found break statement outside of loop")
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

func evalPrefixExpression(operator string, operand ast.Expression, env object.Environment) object.Object {
	evaluatedOperand := Eval(operand, env)
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
	if operand == FALSE || operand == NULL {
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

func evalInfixExpression(operator string, left ast.Expression, right ast.Expression, env object.Environment) object.Object {
	evaluatedLeft := Eval(left, env)
	if isError(evaluatedLeft) {
		return evaluatedLeft
	}

	evaluatedRight := Eval(right, env)
	if isError(evaluatedRight) {
		return evaluatedRight
	}

	switch {
	case evaluatedLeft.Type() == object.INTEGER_OBJ && evaluatedRight.Type() == object.INTEGER_OBJ:
		return evaluateIntegerInfixExpression(operator, evaluatedLeft.(*object.Integer), evaluatedRight.(*object.Integer))
	case evaluatedLeft.Type() == object.STRING_OBJ && evaluatedRight.Type() == object.STRING_OBJ:
		return evaluateStringInfixExpression(operator, evaluatedLeft.(*object.String), evaluatedRight.(*object.String))

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
		return getBoolObject(leftVal == rightVal)
	case "!=":
		return getBoolObject(leftVal != rightVal)
	case "<":
		return getBoolObject(leftVal < rightVal)
	case "<=":
		return getBoolObject(leftVal <= rightVal)
	case ">":
		return getBoolObject(leftVal > rightVal)
	case ">=":
		return getBoolObject(leftVal >= rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evaluateStringInfixExpression(operator string, left *object.String, right *object.String) object.Object {
	leftVal := left.Value
	rightVal := right.Value

	switch operator {
	case "+":
		return &object.String{Value: leftVal + rightVal}
	case "==":
		return getBoolObject(leftVal == rightVal)
	case "!=":
		return getBoolObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalConditionalExpression(condition ast.Expression, consequence *ast.BlockStatement, alternative *ast.BlockStatement, env object.Environment) object.Object {
	evaluatedCondition := Eval(condition, env)
	if isError(evaluatedCondition) {
		return evaluatedCondition
	}

	if isTruthy(evaluatedCondition) {
		return Eval(consequence, env)
	} else if alternative != nil {
		return Eval(alternative, env)
	}

	return NULL
}

func evalBlock(statements []ast.Statement, env object.Environment) object.Object {
	var result object.Object
	for _, stmt := range statements {
		result = Eval(stmt, env)
		if isError(result) || isReturnValue(result) || isBreakStatement(result) {
			return result
		}
	}
	return result
}

func evalLetStatement(name *ast.Identifier, value ast.Expression, env object.Environment) object.Object {
	evaluatedValue := Eval(value, env)
	if isError(evaluatedValue) {
		return evaluatedValue
	}

	env.Set(name.Value, evaluatedValue)

	return NULL
}

func evaluateCallExpression(function ast.Expression, arguments []ast.Expression, env object.Environment) object.Object {
	evaluatedFunc := Eval(function, env)
	if isError(evaluatedFunc) {
		return evaluatedFunc
	}

	if evaluatedFunc.Type() != object.FUNCTION_OBJ {
		return newError("not a function: %s", evaluatedFunc.Type())
	}
	funcObj := evaluatedFunc.(*object.Function)

	evaluatedArgs := evaluateArgExpression(arguments, env)
	if len(evaluatedArgs) == 1 && isError(evaluatedArgs[0]) {
		return evaluatedArgs[0]
	}

	enclosedEnv := extendEnvAndBindArgs(funcObj, evaluatedArgs, env)

	result := Eval(funcObj.Body, enclosedEnv)

	return unwrapReturnValue(result)
}

func evaluateArgExpression(arguments []ast.Expression, env object.Environment) []object.Object {
	var evaluatedArgs []object.Object
	for _, arg := range arguments {
		result := Eval(arg, env)
		if isError(result) {
			return []object.Object{result}
		}
		evaluatedArgs = append(evaluatedArgs, result)
	}
	return evaluatedArgs
}

func extendEnvAndBindArgs(function *object.Function, args []object.Object, env object.Environment) object.Environment {
	enclosedEnv := object.NewEnclosedEnvironment(env)

	for idx, param := range function.Params {
		enclosedEnv.Set(param.Value, args[idx])
	}

	return enclosedEnv
}

func evalWhileLoopExpression(condition ast.Expression, body *ast.BlockStatement, env object.Environment) object.Object {
	var result object.Object = NULL
	for {
		evaluatedCondition := Eval(condition, env)
		if isError(evaluatedCondition) {
			return evaluatedCondition
		}

		isConditionTruthy := isTruthy(evaluatedCondition)
		if !isConditionTruthy {
			break
		}

		result = Eval(body, env)
		if isError(result) || isReturnValue(result) {
			return result
		}

		if isBreakStatement(result) {
			return NULL
		}
	}
	return result
}

func evalAssignmentStatement(name *ast.Identifier, value ast.Expression, env object.Environment) object.Object {
	evaluatedValue := Eval(value, env)
	if isError(evaluatedValue) {
		return evaluatedValue
	}

	env.Set(name.Value, evaluatedValue)

	return NULL
}

func newError(messageTemplate string, args ...any) *object.Error {
	return &object.Error{Message: fmt.Sprintf(messageTemplate, args...)}
}

func isError(obj object.Object) bool {
	return obj != nil && obj.Type() == object.ERROR_OBJ
}

func isTruthy(obj object.Object) bool {
	if obj == FALSE || obj == NULL {
		return false
	}
	return true
}

func isReturnValue(obj object.Object) bool {
	return obj != nil && obj.Type() == object.RETURN_VALUE_OBJ
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}

func isBreakStatement(obj object.Object) bool {
	return obj == BREAK
}
