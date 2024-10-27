package fever

import (
	"disco/base"
)

type PrognEnd struct{}

func NewPrognEnd() BuiltinFeverIF {
	return &PrognEnd{}
}

func init() {
	BuiltinExecutors[base.PROGN_END] = NewPrognEnd()
}

func (p *PrognEnd) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	s := VM.PeekStack()

	VM.RelocationStack()

	VM.PushStack(s)

	return pc, nil
}
