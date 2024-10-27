package fever

import (
	"disco/base"
)

type And struct{}

func NewAnd() BuiltinFeverIF {
	return &And{}
}

func init() {
	BuiltinExecutors[base.AND] = NewAnd()
}

func (a *And) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	b := VM.PeekStack()

	if b.Type != base.NIL {
		pc += 2
		return pc, nil
	}

	return pc, nil
}
