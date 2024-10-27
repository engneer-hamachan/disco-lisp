package fever

import (
	"disco/base"
	"time"
)

type TimeStart struct{}

func NewTimeStart() BuiltinFeverIF {
	return &TimeStart{}
}

func init() {
	BuiltinExecutors[base.TIME_START] = NewTimeStart()
}

func (t *TimeStart) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	VM.Time = time.Now()

	return pc, nil
}
