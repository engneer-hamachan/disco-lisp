package fever

import (
	"disco/base"
	"fmt"
)

type Intern struct{}

func NewIntern() BuiltinFeverIF {
	return &Intern{}
}

func init() {
	BuiltinExecutors[base.INTERN] = NewIntern()
}

func (i *Intern) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	s := VM.PopStack()

	if s.Type != base.STRING {
		return pc, fmt.Errorf("%s is not string", s.Val)
	}

	VM.PushStack(base.MakeSym(s.Val.(string)))

	return pc, nil
}
