package fever

import (
	"disco/base"
	"regexp"
)

type RegexpReplaceAll struct{}

func NewRegexpReplaceAll() BuiltinFeverIF {
	return &RegexpReplaceAll{}
}

func init() {
	BuiltinExecutors[base.REGEXP_REPLACE_ALL] = NewRegexpReplaceAll()
}

func (re *RegexpReplaceAll) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	part_of_stack := VM.PopMultiStack(3)

	target := part_of_stack[0].Val.(string)
	reg_str := part_of_stack[1].Val.(string)
	after_str := part_of_stack[2].Val.(string)

	r := regexp.MustCompile(reg_str)
	new_str := r.ReplaceAllString(target, after_str)

	VM.PushStack(base.MakeString(new_str))

	return pc, nil
}
