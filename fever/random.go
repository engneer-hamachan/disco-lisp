package fever

import (
	"disco/base"
	"math/rand"
)

type Random struct{}

func NewRandom() BuiltinFeverIF {
	return &Random{}
}

func init() {
	BuiltinExecutors[base.RANDOM] = NewRandom()
}

func (r *Random) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	s := VM.PopStack()

	random := rand.Int63n(s.Val.(int64))

	VM.PushStack(base.MakeInt(random))

	return pc, nil
}
