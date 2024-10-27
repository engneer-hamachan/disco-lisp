package fever

import (
	"disco/base"
)

type ListFuntction struct{}

func NewListFuntction() BuiltinFeverIF {
	return &ListFuntction{}
}

func init() {
	BuiltinExecutors[base.LIST_FUNCTION] = NewListFuntction()
}

func (l *ListFuntction) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	pc += 1

	part_of_stack := VM.PopMultiStack(codes[pc].(int))

	s := l.makeList(part_of_stack)

	VM.PushStack(s)

	return pc, nil
}

func (l *ListFuntction) makeList(part_of_stack []*base.S) *base.S {
	if len(part_of_stack) == 0 {
		return base.NilAtom
	}

	return base.Cons(part_of_stack[0], l.makeList(part_of_stack[1:]))
}
