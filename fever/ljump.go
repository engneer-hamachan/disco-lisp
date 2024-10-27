package fever

import (
	"disco/base"
)

type Ljump struct{}

func NewLjump() BuiltinFeverIF {
	return &Ljump{}
}

func init() {
	BuiltinExecutors[base.LJMP] = NewLjump()
}

func (j *Ljump) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	pc += 1

	label := codes[pc]

	for _, code := range codes[pc:] {
		if code == base.LABEL && codes[pc+1] == label {
			break
		}

		pc += 1
	}

	pc += 1

	return pc, nil
}
