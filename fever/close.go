package fever

import (
	"disco/base"
	"os"
)

type Close struct{}

func NewClose() BuiltinFeverIF {
	return &Close{}
}

func init() {
	BuiltinExecutors[base.CLOSE] = NewClose()
}

func (c *Close) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	fp := VM.PopStack().Val.(*os.File)
	fp.Close()

	return pc, nil
}
