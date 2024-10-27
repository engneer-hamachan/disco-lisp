package fever

import (
	"disco/base"
)

type Assoc struct{}

func NewAssoc() BuiltinFeverIF {
	return &Assoc{}
}

func init() {
	BuiltinExecutors[base.ASSOC] = NewAssoc()
}

func (a *Assoc) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	part_of_stack := VM.PopMultiStack(2)

	key := part_of_stack[0]
	target_list := part_of_stack[1]

	for {
		if target_list.GetCar().Type == base.NIL {
			break
		}

		if target_list.GetCaar().Val == key.Val {
			VM.PushStack(target_list.GetCar())
			return pc, nil
		}

		target_list = target_list.GetCdr()
	}

	VM.PushStack(base.NilAtom)

	return pc, nil
}
