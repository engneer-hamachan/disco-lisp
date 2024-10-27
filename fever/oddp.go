package fever

import (
	"disco/base"
)

type Oddp struct{}

func NewOddp() BuiltinFeverIF {
	return &Oddp{}
}

func init() {
	BuiltinExecutors[base.ODDP] = NewOddp()
}

func (o *Oddp) Execute(
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

	if target%2 != 0 {
		VM.PushStack(base.TrueAtom)
		return pc, nil
	}

	VM.PushStack(base.NilAtom)

	return pc, nil
}
