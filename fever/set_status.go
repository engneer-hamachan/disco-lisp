package fever

import (
	"disco/base"
)

type SetStatus struct{}

func NewSetStatus() BuiltinFeverIF {
	return &SetStatus{}
}

func init() {
	BuiltinExecutors[base.SETSTATUS] = NewSetStatus()
}

func (se *SetStatus) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	s := VM.PopStack()

	writer.WriteHeader(int(s.Val.(int64)))

	VM.PushStack(base.TrueAtom)

	return pc, nil
}
