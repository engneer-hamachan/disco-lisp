package fever

import (
	"disco/base"
)

type SetHash struct{}

func NewSetHash() BuiltinFeverIF {
	return &SetHash{}
}

func init() {
	BuiltinExecutors[base.SET_HASH] = NewSetHash()
}

func (sh *SetHash) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	part_of_stack := VM.PopMultiStack(3)

	table := part_of_stack[1].Val.(map[string]*base.S)

	table[part_of_stack[0].Val.(string)] = part_of_stack[2]

	VM.PushStack(part_of_stack[1])

	return pc, nil
}
