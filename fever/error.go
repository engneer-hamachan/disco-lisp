package fever

import (
	"disco/base"
	"fmt"
)

type Error struct{}

func NewError() BuiltinFeverIF {
	return &Error{}
}

func init() {
	BuiltinExecutors[base.ERROR] = NewError()
}

func (e *Error) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	message := VM.PopStack()

	switch message.Val.(type) {
	case string:
		break

	default:
		return pc, fmt.Errorf("argument error! error is want string")
	}

	return pc, fmt.Errorf(message.Val.(string))
}
