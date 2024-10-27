package fever

import (
	"disco/base"
)

type Nth struct{}

func NewNth() BuiltinFeverIF {
	return &Nth{}
}

func init() {
	BuiltinExecutors[base.NTH] = NewNth()
}

func (n *Nth) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	part_of_stack := VM.PopMultiStack(2)

	index := part_of_stack[0].Val.(int64)
	target_s := part_of_stack[1]

	ct := int64(0)

	s := base.NilAtom

	for ct < index && target_s.GetCdr() != nil {
		target_s = target_s.GetCdr()

		ct++
	}

	s = target_s.GetCar()

	VM.PushStack(s)

	return pc, nil
}
