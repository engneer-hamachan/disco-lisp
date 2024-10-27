package fever

import (
	"disco/base"
)

type Add struct{}

func NewAdd() BuiltinFeverIF {
	return &Add{}
}

func init() {
	BuiltinExecutors[base.ADD] = NewAdd()
}

func (a *Add) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	pc += 1

	length_of_args := codes[pc].(int)

	part_of_stack := VM.PopMultiStack(length_of_args)

	answer, err := calculate(part_of_stack, '+')
	if err != nil {
		pc -= 2
		return pc, err
	}

	VM.PushStack(answer)

	return pc, nil
}
