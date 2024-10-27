package fever

import (
	"disco/base"
)

type Lambda struct{}

func NewLambda() BuiltinFeverIF {
	return &Lambda{}
}

func init() {
	BuiltinExecutors[base.LAMBDA] = NewLambda()
}

func (l *Lambda) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	pc += 1

	copy_env := *env

	f := codes[pc].(*base.F)
	f.Env = &copy_env

	s := base.MakeExecFunc(f)
	VM.PushStack(s)

	return pc, nil
}
