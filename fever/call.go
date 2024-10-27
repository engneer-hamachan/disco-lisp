package fever

import (
	"disco/base"
)

type Call struct{}

func NewCall() BuiltinFeverIF {
	return &Call{}
}

func init() {
	BuiltinExecutors[base.CALL] = NewCall()
}

func (c *Call) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	f := VM.PopStack().Val.(*base.F)

	isOptimize := isTailOptimization(codes, caller, pc, f.Name)

	dargs := f.Args

	pc++

	length_of_args := codes[pc].(int)

	switch isOptimize {
	case true:
		setArgs(f.Env, dargs, length_of_args)

		return -1, nil

	default:
		f.Env.PushStack()

		setArgs(f.Env, dargs, length_of_args)

		Fever(f.Body, f.Env, f.Name)

		f.Env.PopStack()

		return pc, nil
	}
}

func isTailOptimization(
	codes []any,
	caller any,
	pc int,
	fname any,
) bool {

	if len(codes[pc:]) == 1 && fname == caller {
		return true
	}

	return false
}

func setArgs(
	env *base.Environment,
	dargs []*base.S,
	length_of_args int,
) error {

	arguments := VM.PopMultiStack(length_of_args)

	env.SetMultiSymbolValueForPeek(dargs, arguments)

	return nil
}
