package fever

import (
	"disco/base"
)

type Label struct{}

func NewLabel() BuiltinFeverIF {
	return &Label{}
}

func init() {
	BuiltinExecutors[base.LABEL] = NewLabel()
}

func (l *Label) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	pc += 1

	return pc, nil
}
