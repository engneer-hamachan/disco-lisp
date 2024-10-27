package fever

import (
	"disco/base"
)

type VectorLen struct{}

func NewVectorLen() BuiltinFeverIF {
	return &VectorLen{}
}

func init() {
	BuiltinExecutors[base.VECTOR_LEN] = NewVectorLen()
}

func (v *VectorLen) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	stack := VM.PopStack()

	vector := stack.Val.([]*base.S)

	VM.PushStack(base.MakeInt(int64(len(vector))))

	return pc, nil
}
