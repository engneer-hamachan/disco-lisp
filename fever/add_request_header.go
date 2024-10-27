package fever

import (
	"disco/base"
	"net/http"
)

type AddRequestHeader struct{}

func NewAddRequestHeader() BuiltinFeverIF {
	return &AddRequestHeader{}
}

func init() {
	BuiltinExecutors[base.ADD_REQUEST_HEADER] = NewAddRequestHeader()
}

func (a *AddRequestHeader) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	part_of_stack := VM.PopMultiStack(3)

	request := part_of_stack[0].Val.(*http.Request)

	request.Header.Set(
		part_of_stack[1].Val.(string),
		part_of_stack[2].Val.(string),
	)

	VM.PushStack(base.MakeReq(request))

	return pc, nil
}
