package fever

import (
	"disco/base"
)

type Type struct{}

func NewType() BuiltinFeverIF {
	return &Type{}
}

func init() {
	BuiltinExecutors[base.TYPE] = NewType()
}

func (p *Type) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	s := VM.PeekStack()

	type_s := base.MakeString(base.TypeToString(s.Type))

	VM.PushStack(type_s)

	return pc, nil
}
