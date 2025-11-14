package object

type Error struct {
	Message string
}

func (e *Error) Type() Type {
	return ERROR_OBJ
}

func (e *Error) Inspect() string {
	return e.Message
}
