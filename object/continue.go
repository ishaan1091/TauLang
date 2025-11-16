package object

type Continue struct {
}

func (c *Continue) Type() Type {
	return CONTINUE_OBJ
}

func (c *Continue) Inspect() string {
	return "continue"
}
