package fever

import (
	"disco/base"
)

type Member struct{}

func NewMember() BuiltinFeverIF {
	return &Member{}
}

func init() {
	BuiltinExecutors[base.MEMBER] = NewMember()
}

func (m *Member) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	part_of_stack := VM.PopMultiStack(2)

	test_atom := part_of_stack[0]
	target_list := part_of_stack[1]

	for target_list.Type != base.NIL {
		if test_atom.Val == target_list.GetCar().Val {
			VM.PushStack(target_list)
			return pc, nil
		}

		target_list = target_list.GetCdr()
	}

	VM.PushStack(base.NilAtom)

	return pc, nil
}
