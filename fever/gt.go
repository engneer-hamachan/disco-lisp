package fever

import (
	"disco/base"
)

type Gt struct{}

func NewGt() BuiltinFeverIF {
	return &Gt{}
}

func init() {
	BuiltinExecutors[base.GT] = NewGt()
}

func (g *Gt) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	cadr := VM.PopStack()
	car := VM.PopStack()

	if car.Val.(int64) > cadr.Val.(int64) {
		VM.PushStack(base.TrueAtom)
		return pc, nil
	}

	VM.PushStack(base.NilAtom)

	return pc, nil
}
