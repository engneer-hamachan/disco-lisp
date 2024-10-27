package fever

import (
	"disco/base"
)

type Stringp struct{}

func NewStringp() BuiltinFeverIF {
	return &Stringp{}
}

func init() {
	BuiltinExecutors[base.STRINGP] = NewStringp()
}

func (st *Stringp) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	s := VM.PopStack()

	if s.Type == base.STRING {
		VM.PushStack(base.TrueAtom)
		return pc, nil
	}

	VM.PushStack(base.NilAtom)

	return pc, nil
}
