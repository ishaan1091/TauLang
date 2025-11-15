package object

type String struct {
	Value string
}

func (s *String) Type() Type {
	return STRING_OBJ
}

func (s *String) Inspect() string {
	return s.Value
}
