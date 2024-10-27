package fever

import (
	"disco/base"
)

type Cdr struct{}

func NewCdr() BuiltinFeverIF {
	return &Cdr{}
}

func init() {
	BuiltinExecutors[base.CDR] = NewCdr()
}

func (c *Cdr) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	s := VM.PopStack()

	VM.PushStack(s.GetCdr())

	return pc, nil
}
