package fever

import (
	"disco/base"
)

type VectorPush struct{}

func NewVectorPush() BuiltinFeverIF {
	return &VectorPush{}
}

func init() {
	BuiltinExecutors[base.VECTOR_PUSH] = NewVectorPush()
}

func (v *VectorPush) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	part_of_stack := VM.PopMultiStack(2)

	vector := part_of_stack[0].Val.([]*base.S)
	vector = append(vector, part_of_stack[1])

	vector_s := base.MakeVector(vector)

	VM.PushStack(vector_s)

	return pc, nil
}
