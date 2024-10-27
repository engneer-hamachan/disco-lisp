package fever

import (
	"disco/base"
)

type Ge struct{}

func NewGe() BuiltinFeverIF {
	return &Ge{}
}

func init() {
	BuiltinExecutors[base.GE] = NewGe()
}

func (g *Ge) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	cadr := VM.PopStack()
	car := VM.PopStack()

	if car.Val.(int64) >= cadr.Val.(int64) {
		VM.PushStack(base.TrueAtom)
		return pc, nil
	}

	VM.PushStack(base.NilAtom)

	return pc, nil
}
