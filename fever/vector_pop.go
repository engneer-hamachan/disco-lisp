package fever

import (
	"disco/base"
)

type VectorPop struct{}

func NewVectorPop() BuiltinFeverIF {
	return &VectorPop{}
}

func init() {
	BuiltinExecutors[base.VECTOR_POP] = NewVectorPop()
}

func (v *VectorPop) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	stack := VM.PopStack()

	vector := stack.Val.([]*base.S)
	tail_idx := len(vector) - 1

	stack.Val = vector[:tail_idx]
	vector_s := vector[tail_idx]

	VM.PushStack(vector_s)

	return pc, nil
}
