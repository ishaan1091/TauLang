package object

type Type string

const (
	INTEGER_OBJ = "INTEGER"
)

type Object interface {
	Type() Type
	Inspect() string
}
