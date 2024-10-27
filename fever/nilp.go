package fever

import (
	"disco/base"
)

type Nilp struct{}

func NewNilp() BuiltinFeverIF {
	return &Nilp{}
}

func init() {
	BuiltinExecutors[base.NILP] = NewNilp()
}

func (n *Nilp) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	s := VM.PopStack()

	if s.Type == base.NIL {
		VM.PushStack(base.TrueAtom)
		return pc, nil
	}

	VM.PushStack(base.NilAtom)

	return pc, nil
}
