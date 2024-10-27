package fever

import (
	"disco/base"
)

type SubString struct{}

func NewSubString() BuiltinFeverIF {
	return &SubString{}
}

func init() {
	BuiltinExecutors[base.SUBSTRING] = NewSubString()
}

func (sp *SubString) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	part_of_stack := VM.PopMultiStack(3)

	target := part_of_stack[0].Val.(string)
	start_idx := part_of_stack[1].Val.(int64)
	end_idx := part_of_stack[2].Val.(int64)
	length := start_idx + end_idx
	max := int64(len(target))

	if length > max {
		length = max
	}

	VM.PushStack(base.MakeString(target[start_idx:length]))

	return pc, nil
}
