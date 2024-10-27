package fever

import (
	"bytes"
	"disco/base"
	"encoding/json"
	"fmt"
	"net/http"
)

type MakeRequest struct{}

func NewMakeRequest() BuiltinFeverIF {
	return &MakeRequest{}
}

func init() {
	BuiltinExecutors[base.MAKE_REQUEST] = NewMakeRequest()
}

func (m *MakeRequest) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	pc += 1

	count_of_args := codes[pc].(int)

	part_of_stack := VM.PopMultiStack(count_of_args)

	switch count_of_args {
	case 2:
		request, err :=
			http.NewRequest(
				part_of_stack[0].Val.(string),
				part_of_stack[1].Val.(string),
				nil,
			)
		if err != nil {
			return pc, err
		}

		VM.PushStack(base.MakeReq(request))

	case 3:
		body := part_of_stack[2].Val.(string)

		request, err :=
			m.makeRequestWithBody(
				part_of_stack[0].Val.(string),
				part_of_stack[1].Val.(string),
				body,
			)

		if err != nil {
			return pc, err
		}

		VM.PushStack(base.MakeReq(request))

	default:
		return pc, fmt.Errorf("make-request argument error!")
	}

	return pc, nil
}

func (m *MakeRequest) makeRequestWithBody(
	method string,
	url string,
	body string,
) (*http.Request, error) {

	var request *http.Request
	var err error

	switch m.isJson(body) {
	case true:
		request_body := bytes.NewBuffer([]byte(body))

		request, err =
			http.NewRequest(
				method,
				url,
				request_body,
			)
		if err != nil {
			return nil, err
		}

		request.Header.Set("Content-Type", "application/json")

	default:
		request_body := bytes.NewBufferString(body)

		request, err =
			http.NewRequest(
				method,
				url,
				request_body,
			)
		if err != nil {
			return nil, err
		}

		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	return request, nil
}

func (m *MakeRequest) isJson(str string) bool {
	var json_raw_msg json.RawMessage
	err := json.Unmarshal([]byte(str), &json_raw_msg)

	return err == nil
}
