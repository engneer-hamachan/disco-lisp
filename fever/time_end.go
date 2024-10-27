package fever

import (
	"disco/base"
	"fmt"
	"time"
)

type TimeEnd struct{}

func NewTimeEnd() BuiltinFeverIF {
	return &TimeEnd{}
}

func init() {
	BuiltinExecutors[base.TIME_END] = NewTimeEnd()
}

func (t *TimeEnd) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	fmt.Println(time.Now().Sub(VM.Time))

	return pc, nil
}
