package fever

import (
	"disco/base"
	"os"
)

type ReadFile struct{}

func NewReadFile() BuiltinFeverIF {
	return &ReadFile{}
}

func init() {
	BuiltinExecutors[base.READ_FILE] = NewReadFile()
}

func (r *ReadFile) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	fp := VM.PopStack().Val.(*os.File)
	defer fp.Close()

	str := make([]byte, 4096)

	byte_count, err := fp.Read(str)
	if err != nil {
		return pc, err
	}

	s := base.MakeString(string(str[:byte_count]))

	VM.PushStack(s)

	return pc, nil
}
