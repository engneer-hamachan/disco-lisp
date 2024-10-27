package fever

import (
	"disco/base"
	"net/http"
)

type DefServer struct{}

var server map[string]*http.Server = make(map[string]*http.Server)

func NewDefServer() BuiltinFeverIF {
	return &DefServer{}
}

func init() {
	BuiltinExecutors[base.DEFSERVER] = NewDefServer()
}

func (d *DefServer) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	part_of_stack := VM.PopMultiStack(2)

	server[part_of_stack[0].Val.(string)] =
		&http.Server{
			Addr:    part_of_stack[1].Val.(string),
			Handler: nil,
		}

	VM.PushStack(base.TrueAtom)

	return pc, nil
}
