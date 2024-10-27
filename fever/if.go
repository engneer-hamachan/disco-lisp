package fever

import (
	"disco/base"
)

type If struct{}

func NewIf() BuiltinFeverIF {
	return &If{}
}

func init() {
	BuiltinExecutors[base.IF] = NewIf()
}

func (i *If) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	pc += 1

	jump_count := codes[pc].(int)

	b := VM.PopStack()

	if b.Type == base.NIL {
		pc += jump_count
		return pc, nil
	}

	return pc, nil
}
