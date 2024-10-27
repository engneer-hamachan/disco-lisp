package fever

import (
	"disco/base"
)

type Pairp struct{}

func NewPairp() BuiltinFeverIF {
	return &Pairp{}
}

func init() {
	BuiltinExecutors[base.PAIRP] = NewPairp()
}

func (pa *Pairp) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	s := VM.PopStack()

	if s.GetCdr().Type != base.LIST && s.GetCdr().Type != base.NIL {
		VM.PushStack(base.TrueAtom)
		return pc, nil
	}

	VM.PushStack(base.NilAtom)

	return pc, nil
}
