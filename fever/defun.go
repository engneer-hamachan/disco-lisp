package fever

import (
	"disco/base"
)

type Defun struct{}

var DefineFunctions = make(map[string]*base.S)

func NewDefun() BuiltinFeverIF {
	return &Defun{}
}

func init() {
	BuiltinExecutors[base.DEFUN] = NewDefun()
}

func (d *Defun) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	pc += 1

	f := codes[pc].(*base.F)
	copy_env := *env

	f.Env = &copy_env
	DefineFunctions[f.Name.(string)] = base.MakeExecFunc(f)

	VM.PushStack(base.TrueAtom)

	return pc, nil
}
