package fever

import (
	"disco/base"
)

type Evenp struct{}

func NewEvenp() BuiltinFeverIF {
	return &Evenp{}
}

func init() {
	BuiltinExecutors[base.EVENP] = NewEvenp()
}

func (e *Evenp) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	target := VM.PopStack().Val.(int64)

	if target == 0 {
		VM.PushStack(base.NilAtom)
		return pc, nil
	}

	if target%2 == 0 {
		VM.PushStack(base.TrueAtom)
		return pc, nil
	}

	VM.PushStack(base.NilAtom)

	return pc, nil
}
