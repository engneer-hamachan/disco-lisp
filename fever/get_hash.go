package fever

import (
	"disco/base"
)

type GetHash struct{}

func NewGetHash() BuiltinFeverIF {
	return &GetHash{}
}

func init() {
	BuiltinExecutors[base.GET_HASH] = NewGetHash()
}

func (g *GetHash) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	part_of_stack := VM.PopMultiStack(2)

	table := part_of_stack[0].Val.(map[string]*base.S)

	s, ok := table[part_of_stack[1].Val.(string)]
	if ok {
		VM.PushStack(s)
		return pc, nil
	}

	VM.PushStack(base.NilAtom)
	return pc, nil
}
