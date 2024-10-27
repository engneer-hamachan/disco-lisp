package fever

import (
	"disco/base"
	"fmt"
	"os"
)

type Open struct{}

func NewOpen() BuiltinFeverIF {
	return &Open{}
}

func init() {
	BuiltinExecutors[base.OPEN] = NewOpen()
}

func (o *Open) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	part_of_stack := VM.PopMultiStack(2)

	path := part_of_stack[0]
	mode := part_of_stack[1]

	var fp *os.File
	var err error

	switch mode.Val.(string) {
	case "r":
		fp, err = os.Open(path.Val.(string))
		if err != nil {
			return pc, err
		}

	case "w":
		fp, err = os.Create(path.Val.(string))
		if err != nil {
			return pc, err
		}

	default:
		return pc, fmt.Errorf("open mode is w (write) or r (read)")
	}

	VM.PushStack(base.MakeFp(fp))

	return pc, nil
}
