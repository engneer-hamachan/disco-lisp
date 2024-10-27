package fever

import (
	"disco/base"
	"os"
)

type Write struct{}

func NewWrite() BuiltinFeverIF {
	return &Write{}
}

func init() {
	BuiltinExecutors[base.WRITE] = NewWrite()
}

func (w *Write) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	part_of_stack := VM.PopMultiStack(2)

	fp := part_of_stack[0].Val.(*os.File)
	defer fp.Close()

	str := part_of_stack[1].Val.(string)

	_, err := fp.Write([]byte(str))
	if err != nil {
		return pc, err
	}

	return pc, nil
}
