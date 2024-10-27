package fever

import (
	"disco/base"
	"io"
	"net/http"
)

type DoRequest struct{}

func NewDoRequest() BuiltinFeverIF {
	return &DoRequest{}
}

func init() {
	BuiltinExecutors[base.DO_REQUEST] = NewDoRequest()
}

func (d *DoRequest) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	s := VM.PopStack()

	http_client := new(http.Client)

	response, err := http_client.Do(s.Val.(*http.Request))
	if err != nil {
		return pc, err
	}

	defer response.Body.Close()

	byteArray, err := io.ReadAll(response.Body)
	if err != nil {
		return pc, err
	}

	VM.PushStack(base.MakeString(string(byteArray)))

	return pc, nil
}
