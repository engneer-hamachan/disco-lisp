package fever

import (
	"disco/base"
)

type Global struct{}

func NewGlobal() BuiltinFeverIF {
	return &Global{}
}

func init() {
	BuiltinExecutors[base.GLOBAL] = NewGlobal()
}

func (g *Global) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	part_of_stack := VM.PopMultiStack(2)

	car := part_of_stack[0]
	cadr := part_of_stack[1]

	if car.Type == base.SYMBOL {
		switch cadr.Type {
		case base.FUNC:
			f := DefineFunctions[cadr.Val.(string)]
			env.SetGlobalSymbolValue(car, f)

		default:
			env.SetGlobalSymbolValue(car, cadr)
		}
	}

	VM.PushStack(cadr)

	return pc, nil
}
