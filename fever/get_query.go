package fever

import (
	"disco/base"
)

type GetQuery struct{}

func NewGetQuery() BuiltinFeverIF {
	return &GetQuery{}
}

func init() {
	BuiltinExecutors[base.GETQUERY] = NewGetQuery()
}

func (d *GetQuery) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	s := VM.PopStack()

	v, ok := queries[s.Val.(string)]
	if !ok {
		VM.PushStack(base.NilAtom)
		return pc, nil
	}

	VM.PushStack(v)

	return pc, nil
}
