package fever

import (
	"disco/base"
)

type Or struct{}

func NewOr() BuiltinFeverIF {
	return &Or{}
}

func init() {
	BuiltinExecutors[base.OR] = NewOr()
}

func (a *Or) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	b := VM.PeekStack()

	if b.Type == base.NIL {
		pc += 2
	}

	return pc, nil
}
