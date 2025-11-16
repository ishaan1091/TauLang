package object

type Environment interface {
	Get(key string) (Object, bool)
	Set(key string, value Object) Object
}

type environment struct {
	store    map[string]Object
	outerEnv Environment
}

func NewEnclosedEnvironment(outerEnv Environment) Environment {
	return &environment{
		store:    map[string]Object{},
		outerEnv: outerEnv,
	}
}

func NewEnvironment() Environment {
	return &environment{
		store: map[string]Object{},
	}
}

func (e *environment) Get(key string) (Object, bool) {
	obj, ok := e.store[key]
	if !ok && e.outerEnv != nil {
		return e.outerEnv.Get(key)
	}
	return obj, ok
}

func (e *environment) Set(key string, value Object) Object {
	e.store[key] = value
	return value
}
