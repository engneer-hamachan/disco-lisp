package fever

import (
	"disco/base"
)

type Not struct{}

func NewNot() BuiltinFeverIF {
	return &Not{}
}

func init() {
	BuiltinExecutors[base.NOT] = NewNot()
}

func (n *Not) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	b := VM.PopStack()

	switch b.Type {
	case base.NIL:
		VM.PushStack(base.TrueAtom)

	default:
		VM.PushStack(base.NilAtom)
	}

	return pc, nil
}
