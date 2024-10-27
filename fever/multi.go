package fever

import (
	"disco/base"
)

type Multi struct{}

func NewMulti() BuiltinFeverIF {
	return &Multi{}
}

func init() {
	BuiltinExecutors[base.MULTI] = NewMulti()
}

func (m *Multi) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	pc += 1

	length_of_args := codes[pc].(int)
	part_of_stack := VM.PopMultiStack(length_of_args)

	answer, err := calculate(part_of_stack, '*')
	if err != nil {
		return pc, err
	}

	VM.PushStack(answer)

	return pc, nil
}
