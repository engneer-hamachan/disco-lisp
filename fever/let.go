package fever

import (
	"disco/base"
)

type Let struct{}

func NewLet() BuiltinFeverIF {
	return &Let{}
}

func init() {
	BuiltinExecutors[base.LET] = NewLet()
}

func (se *Let) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	pc += 1

	sym := codes[pc].(string)
	args := VM.PopStack()

	switch args.Type {
	case base.FUNC:
		f := DefineFunctions[args.Val.(string)]
		env.SetSymbolValue(base.MakeSym(sym), f)

	default:
		env.SetSymbolValue(base.MakeSym(sym), args)
	}

	return pc, nil
}
