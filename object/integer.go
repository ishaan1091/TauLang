package object

import "fmt"

type Integer struct {
	Value int64
}

func (i *Integer) Type() Type {
	return INTEGER_OBJ
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Hash() HashKey {
	return HashKey{ObjectType: INTEGER_OBJ, Value: uint64(i.Value)}
}
