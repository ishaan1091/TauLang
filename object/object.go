package object

type Type string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	ERROR_OBJ        = "ERROR"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	STRING_OBJ       = "STRING"
	FUNCTION_OBJ     = "FUNCTION"
	BREAK_OBJ        = "BREAK"
	CONTINUE_OBJ     = "CONTINUE"
	BUILTIN_OBJ      = "BUILTIN"
)

type Object interface {
	Type() Type
	Inspect() string
}
