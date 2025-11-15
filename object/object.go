package object

type Type string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	ERROR_OBJ   = "ERROR"
	NULL_OBJ    = "NULL"
)

type Object interface {
	Type() Type
	Inspect() string
}
