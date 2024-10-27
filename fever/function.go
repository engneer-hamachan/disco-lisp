package fever

import (
	"disco/base"
	"fmt"
)

type Function struct{}

func NewFunction() BuiltinFeverIF {
	return &Function{}
}

func init() {
	BuiltinExecutors[base.FUNCTION] = NewFunction()
}

func (fu *Function) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	s := VM.PopStack().Val

	switch s.(type) {
	case *base.F:
		VM.PushStack(base.MakeExecFunc(s.(*base.F)))

	default:
		return pc, fmt.Errorf("%s is not function", s)
	}

	return pc, nil
}
