package fever

import (
	"disco/base"
	"fmt"
)

type Ldf struct{}

func NewLdf() BuiltinFeverIF {
	return &Ldf{}
}

func init() {
	BuiltinExecutors[base.LDF] = NewLdf()
}

func (l *Ldf) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	pc += 1
	fname := codes[pc].(string)

	f, ok := DefineFunctions[fname]
	if ok {
		VM.PushStack(f)
		return pc, nil
	}

	return pc, fmt.Errorf("undefined function %s", fname)
}
