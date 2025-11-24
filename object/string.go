package object

import "hash/fnv"

type String struct {
	Value string
}

func (s *String) Type() Type {
	return STRING_OBJ
}

func (s *String) Inspect() string {
	return s.Value
}

func (s *String) Hash() HashKey {
	return HashKey{ObjectType: STRING_OBJ, Value: hash(s.Value)}
}

func hash(s string) uint64 {
	h := fnv.New64a()
	// Write on fnv.New64a never returns an error because it only operates on memory;
	// the error is part of the io.Writer interface signature. Safe to ignore.
	h.Write([]byte(s))
	return h.Sum64()
}
