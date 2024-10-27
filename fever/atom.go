package fever

import (
	"disco/base"
	"disco/predicater"
)

type Atomp struct{}

func NewAtomp() BuiltinFeverIF {
	return &Atomp{}
}

func init() {
	BuiltinExecutors[base.ATOMP] = NewAtomp()
}

func (a *Atomp) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	s := VM.PopStack()

	if predicater.Atomp(s) {
		VM.PushStack(base.TrueAtom)
		return pc, nil
	}

	VM.PushStack(base.NilAtom)

	return pc, nil
}
