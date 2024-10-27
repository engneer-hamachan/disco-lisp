package fever

import (
	"disco/base"
	"regexp"
)

type RegexpMatch struct{}

func NewRegexpMatch() BuiltinFeverIF {
	return &RegexpMatch{}
}

func init() {
	BuiltinExecutors[base.REGEXP_MATCH] = NewRegexpMatch()
}

func (re *RegexpMatch) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	part_of_stack := VM.PopMultiStack(2)

	target := part_of_stack[0].Val.(string)
	reg_str := part_of_stack[1].Val.(string)

	r := regexp.MustCompile(reg_str)
	b := r.MatchString(target)

	if b {
		VM.PushStack(base.TrueAtom)
		return pc, nil
	}

	VM.PushStack(base.NilAtom)
	return pc, nil
}
