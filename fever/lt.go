package fever

import (
	"disco/base"
)

type Lt struct{}

func NewLt() BuiltinFeverIF {
	return &Lt{}
}

func init() {
	BuiltinExecutors[base.LT] = NewLt()
}

func (g *Lt) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	cadr := VM.PopStack()
	car := VM.PopStack()

	if car.Val.(int64) < cadr.Val.(int64) {
		VM.PushStack(base.TrueAtom)
		return pc, nil
	}

	VM.PushStack(base.NilAtom)

	return pc, nil
}
