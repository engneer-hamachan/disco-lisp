package fever

import (
	"disco/base"
)

type Aref struct{}

func NewAref() BuiltinFeverIF {
	return &Aref{}
}

func init() {
	BuiltinExecutors[base.AREF] = NewAref()
}

func (n *Aref) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	part_of_stack := VM.PopMultiStack(2)

	index := part_of_stack[0].Val.(int64)
	target_vector := part_of_stack[1].Val.([]*base.S)
	target_vector_length := len(target_vector)

	if target_vector_length == 0 || target_vector_length < int(index)+1 {
		VM.PushStack(base.NilAtom)
		return pc, nil
	}

	VM.PushStack(target_vector[index])

	return pc, nil
}
