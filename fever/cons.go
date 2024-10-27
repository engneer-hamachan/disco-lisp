package fever

import (
	"disco/base"
)

type Cons struct{}

func NewCons() BuiltinFeverIF {
	return &Cons{}
}

func init() {
	BuiltinExecutors[base.CONS] = NewCons()
}

func (c *Cons) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	list := VM.PopStack()
	s := VM.PopStack()

	consed_s := base.Cons(s, list)

	VM.PushStack(consed_s)

	return pc, nil
}
