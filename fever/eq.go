package fever

import (
	"disco/base"
)

type Eq struct{}

func NewEq() BuiltinFeverIF {
	return &Eq{}
}

func init() {
	BuiltinExecutors[base.EQ] = NewEq()
}

func (e *Eq) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	car := VM.PopStack()
	cadr := VM.PopStack()

	if car.Val == cadr.Val && car.Type == cadr.Type {
		VM.PushStack(base.TrueAtom)
		return pc, nil
	}

	VM.PushStack(base.NilAtom)

	return pc, nil
}
