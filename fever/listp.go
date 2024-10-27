package fever

import (
	"disco/base"
)

type Listp struct{}

func NewListp() BuiltinFeverIF {
	return &Listp{}
}

func init() {
	BuiltinExecutors[base.LISTP] = NewListp()
}

func (l *Listp) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	s := VM.PopStack()

	if s.Type == base.LIST || s.Type == base.NIL {
		VM.PushStack(base.TrueAtom)
		return pc, nil
	}

	VM.PushStack(base.NilAtom)

	return pc, nil
}
