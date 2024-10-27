package fever

import (
	"disco/base"
)

type Div struct{}

func NewDiv() BuiltinFeverIF {
	return &Div{}
}

func init() {
	BuiltinExecutors[base.DIV] = NewDiv()
}

func (d *Div) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	pc += 1

	length_of_args := codes[pc].(int)

	part_of_stack := VM.PopMultiStack(length_of_args)

	answer, err := calculate(part_of_stack, '/')
	if err != nil {
		return pc, err
	}

	VM.PushStack(answer)

	return pc, nil
}
