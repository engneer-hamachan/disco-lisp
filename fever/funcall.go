package fever

import (
	"disco/base"
	"fmt"
)

type Funcall struct{}

func NewFuncall() BuiltinFeverIF {
	return &Funcall{}
}

func init() {
	BuiltinExecutors[base.FUNCALL] = NewFuncall()
}

func (c *Funcall) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	var f *base.F

	s := VM.PopStack().Val

	switch s.(type) {
	case *base.F:
		f = s.(*base.F)
	default:
		return pc, fmt.Errorf("not function args.")
	}

	pc += 1

	args_length := codes[pc].(int)
	dargs_length := len(f.Args)

	if args_length > dargs_length {
		return pc, fmt.Errorf(
			"too many arguments %s want %d arguments",
			f.Name,
			dargs_length,
		)
	}

	if args_length < dargs_length {
		return pc, fmt.Errorf(
			"too few arguments %s want %d arguments",
			f.Name,
			dargs_length,
		)
	}

	dargs := f.Args

	switch isTailOptimization(codes, caller, pc, f.Name) {
	case true:
		setArgs(f.Env, dargs, args_length)

		return -1, nil

	default:
		f.Env.PushStack()

		setArgs(f.Env, dargs, args_length)

		Fever(f.Body, f.Env, f.Name)

		f.Env.PopStack()

		return pc, nil
	}
}
