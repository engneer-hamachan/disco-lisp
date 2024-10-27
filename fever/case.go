package fever

import (
	"disco/base"
)

type Case struct{}

func NewCase() BuiltinFeverIF {
	return &Case{}
}

func init() {
	BuiltinExecutors[base.CASE] = NewCase()
}

func (c *Case) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	pc += 1

	jump_count := codes[pc].(int)

	part_of_stack := VM.PopMultiStack(2)

	switch part_of_stack[1].Val {
	case part_of_stack[0].Val, "default":
		break

	default:
		pc += jump_count
	}

	return pc, nil
}
