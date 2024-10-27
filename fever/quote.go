package fever

import (
	"disco/base"
)

type Quote struct{}

func NewQuote() BuiltinFeverIF {
	return &Quote{}
}

func init() {
	BuiltinExecutors[base.QUOTE] = NewQuote()
}

func (q *Quote) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	pc += 1

	sym := codes[pc]

	VM.PushStack(sym.(*base.S))

	return pc, nil
}
