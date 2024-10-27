package fever

import (
	"disco/base"
)

type RunServer struct{}

func NewRunServer() BuiltinFeverIF {
	return &RunServer{}
}

func init() {
	BuiltinExecutors[base.RUNSERVER] = NewRunServer()
}

func (r *RunServer) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	s := VM.PopStack()

	server[s.Val.(string)].ListenAndServe()

	VM.PushStack(base.TrueAtom)

	return pc, nil
}
