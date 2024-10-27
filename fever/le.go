package fever

import (
	"disco/base"
)

type Le struct{}

func NewLe() BuiltinFeverIF {
	return &Le{}
}

func init() {
	BuiltinExecutors[base.LE] = NewLe()
}

func (g *Le) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	cadr := VM.PopStack()
	car := VM.PopStack()

	if car.Val.(int64) <= cadr.Val.(int64) {
		VM.PushStack(base.TrueAtom)
		return pc, nil
	}

	VM.PushStack(base.NilAtom)

	return pc, nil
}
