package fever

import (
	"disco/base"
)

type Jump struct{}

func NewJump() BuiltinFeverIF {
	return &Jump{}
}

func init() {
	BuiltinExecutors[base.JMP] = NewJump()
}

func (j *Jump) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	pc += 1

	jump_count := codes[pc].(int)

	pc += jump_count

	return pc, nil
}
