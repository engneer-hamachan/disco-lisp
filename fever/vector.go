package fever

import (
	"disco/base"
)

type Vector struct{}

func NewVector() BuiltinFeverIF {
	return &Vector{}
}

func init() {
	BuiltinExecutors[base.MAKE_VECTOR] = NewVector()
}

func (v *Vector) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	pc += 1

	length_of_args := codes[pc].(int)
	part_of_stack := VM.PopMultiStack(length_of_args)

	var vector []*base.S

	for _, vec := range part_of_stack {
		vector = append(vector, vec)
	}

	vector_s := base.MakeVector(vector)

	VM.PushStack(vector_s)

	return pc, nil
}
