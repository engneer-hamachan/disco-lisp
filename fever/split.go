package fever

import (
	"disco/base"
	"strings"
)

type Split struct{}

func NewSplit() BuiltinFeverIF {
	return &Split{}
}

func init() {
	BuiltinExecutors[base.SPLIT] = NewSplit()
}

func (sp *Split) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	part_of_stack := VM.PopMultiStack(2)

	separate_char := part_of_stack[0].Val.(string)
	target := part_of_stack[1].Val.(string)

	str_arr := strings.Split(target, separate_char)

	s := sp.arrayToSList(str_arr)

	VM.PushStack(s)

	return pc, nil
}

func (sp *Split) arrayToSList(str_arr []string) *base.S {
	if len(str_arr) < 1 {
		return base.NilAtom
	}

	s := base.Cons(base.MakeString(str_arr[0]), sp.arrayToSList(str_arr[1:]))

	return s
}
