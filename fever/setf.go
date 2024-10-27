package fever

import (
	"disco/base"
)

type Setf struct{}

func NewSetf() BuiltinFeverIF {
	return &Setf{}
}

func init() {
	BuiltinExecutors[base.SET] = NewSetf()
}

func (se *Setf) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	part_of_stack := VM.PopMultiStack(2)

	car := part_of_stack[0]
	cadr := part_of_stack[1]

	switch car.Type {
	case base.SYMBOL:
		if cadr.Type == base.FUNC {
			f := DefineFunctions[cadr.Val.(string)]
			env.SetSymbolValue(car, f)
			break
		}

		env.SetSymbolValue(car, cadr)

	default:
		car.Val = cadr.Val
		car.Type = cadr.Type
	}

	VM.PushStack(cadr)

	return pc, nil
}
