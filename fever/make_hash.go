package fever

import (
	"disco/base"
)

type MakeHash struct{}

func NewMakeHash() BuiltinFeverIF {
	return &MakeHash{}
}

func init() {
	BuiltinExecutors[base.MAKE_HASH] = NewMakeHash()
}

func (m *MakeHash) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	table := make(map[string]*base.S)

	VM.PushStack(base.MakeHash(table))

	return pc, nil
}
