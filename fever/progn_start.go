package fever

import (
	"disco/base"
)

type PrognStart struct{}

func NewPrognStart() BuiltinFeverIF {
	return &PrognStart{}
}

func init() {
	BuiltinExecutors[base.PROGN_START] = NewPrognStart()
}

func (p *PrognStart) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	VM.EvacutionStack()

	return pc, nil
}
