package object

type Break struct {
}

func (b *Break) Type() Type {
	return BREAK_OBJ
}

func (b *Break) Inspect() string {
	return "break"
}
