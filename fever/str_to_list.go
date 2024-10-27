package fever

import (
	"disco/base"
)

type StrToList struct{}

func NewStrToList() BuiltinFeverIF {
	return &StrToList{}
}

func init() {
	BuiltinExecutors[base.STR_TO_LIST] = NewStrToList()
}

func (stl *StrToList) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	str_s := VM.PopStack()

	s := stl.strToSList(str_s.Val.(string))

	VM.PushStack(s)

	return pc, nil
}

func (stl *StrToList) strToSList(str string) *base.S {
	if len(str) < 1 {
		return base.NilAtom
	}

	s := base.Cons(base.MakeString(string(str[0])), stl.strToSList(str[1:]))

	return s
}
