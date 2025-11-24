package evaluator

import (
	"fmt"
	"taulang/ast"
	"taulang/object"
)

var (
	NULL     = &object.Null{}
	BREAK    = &object.Break{}
	CONTINUE = &object.Continue{}
	TRUE     = &object.Boolean{Value: true}
	FALSE    = &object.Boolean{Value: false}
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
		return evalCallExpression(node.Function, node.Arguments, env)
	case *ast.AssignmentStatement:
		return evalAssignmentStatement(node.Name, node.Value, env)
	case *ast.WhileLoopExpression:
		return evalWhileLoopExpression(node.Condition, node.Body, env)
	case *ast.BreakStatement:
		return BREAK
	case *ast.ContinueStatement:
		return CONTINUE
	case *ast.ArrayLiteral:
		return evalArrayLiteral(node.Elements, env)
	case *ast.IndexExpression:
		return evalIndexExpression(node.IndexedExpression, node.Index, env)
	case *ast.HashLiteral:
		return evalHashLiteral(node.Pairs, env)
	default:
		return newError("no defined evaluations for input: %s", node.String())
	}
}

func evalIdentifier(identifierName string, env object.Environment) object.Object {
	if obj, ok := env.Get(identifierName); ok {
		return obj
	}

	if obj, ok := builtins[identifierName]; ok {
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

		if isBreak(result) {
			return newError("found break statement outside of loop")
		}

		if isContinue(result) {
			return newError("found continue statement outside of loop")
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
		if isError(result) || isReturnValue(result) || isBreak(result) || isContinue(result) {
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

func evalCallExpression(function ast.Expression, arguments []ast.Expression, env object.Environment) object.Object {
	evaluatedFunc := Eval(function, env)
	if isError(evaluatedFunc) {
		return evaluatedFunc
	}

	evaluatedArgs := evaluateExpression(arguments, env)
	if len(evaluatedArgs) == 1 && isError(evaluatedArgs[0]) {
		return evaluatedArgs[0]
	}

	switch funcObj := evaluatedFunc.(type) {
	case *object.Function:
		enclosedEnv := extendEnvAndBindArgs(funcObj, evaluatedArgs, env)
		result := Eval(funcObj.Body, enclosedEnv)
		return unwrapReturnValue(result)
	case *object.Builtin:
		return funcObj.Fn(evaluatedArgs...)
	default:
		return newError("not a function: %s", evaluatedFunc.Type())
	}
}

func evaluateExpression(expressions []ast.Expression, env object.Environment) []object.Object {
	var evaluatedArgs []object.Object
	for _, arg := range expressions {
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

		if isBreak(result) {
			return NULL
		}

		if isContinue(result) {
			result = NULL
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

func evalArrayLiteral(elements []ast.Expression, env object.Environment) object.Object {
	evaluatedElements := evaluateExpression(elements, env)
	if len(elements) == 1 && isError(evaluatedElements[0]) {
		return evaluatedElements[0]
	}

	return &object.Array{Elements: evaluatedElements}
}

func evalIndexExpression(expression ast.Expression, index ast.Expression, env object.Environment) object.Object {
	evaluatedIndexedObject := Eval(expression, env)
	if isError(evaluatedIndexedObject) {
		return evaluatedIndexedObject
	}

	evaluatedIndex := Eval(index, env)
	if isError(evaluatedIndex) {
		return evaluatedIndex
	}

	switch {
	case evaluatedIndexedObject.Type() == object.ARRAY_OBJ && evaluatedIndex.Type() == object.INTEGER_OBJ:
		return evalArrayIndexExpression(evaluatedIndexedObject, evaluatedIndex)
	case evaluatedIndexedObject.Type() == object.HASHMAP_OBJ:
		return evalHashIndexExpression(evaluatedIndexedObject, evaluatedIndex)
	default:
		return newError("index operator not supported: %s[%s]", evaluatedIndexedObject.Type(), evaluatedIndex.Type())
	}
}

func evalArrayIndexExpression(indexedObject object.Object, index object.Object) object.Object {
	array := indexedObject.(*object.Array).Elements
	indexVal := index.(*object.Integer).Value

	if indexVal < 0 || indexVal >= int64(len(array)) {
		return NULL
	}

	return array[indexVal]
}

func evalHashIndexExpression(indexedObject object.Object, index object.Object) object.Object {
	hashObject := indexedObject.(*object.HashMap)

	key, ok := index.(object.Hashable)
	if !ok {
		return newError("unusable as hash key: %s", index.Type())
	}

	pair, ok := hashObject.Pairs[key.Hash()]
	if !ok {
		return NULL
	}

	return pair.Value
}

func evalHashLiteral(pairs []ast.HashPair, env object.Environment) object.Object {
	p := make(map[object.HashKey]object.HashPair)

	for _, pair := range pairs {
		keyNode := pair.Key
		valueNode := pair.Value

		key := Eval(keyNode, env)
		if isError(key) {
			return key
		}

		hashKey, ok := key.(object.Hashable)
		if !ok {
			return newError("unusable as hash key: %s", key.Type())
		}

		value := Eval(valueNode, env)
		if isError(value) {
			return value
		}

		hashed := hashKey.Hash()
		p[hashed] = object.HashPair{Key: key, Value: value}
	}

	return &object.HashMap{Pairs: p}
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

func isBreak(obj object.Object) bool {
	return obj == BREAK
}

func isContinue(obj object.Object) bool {
	return obj == CONTINUE
}
