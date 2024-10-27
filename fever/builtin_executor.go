package fever

import (
	"disco/base"
)

type BuiltinFeverIF interface {
	Execute(
		codes []any,
		pc int,
		env *base.Environment,
		caller any,
	) (int, error)
}

var BuiltinExecutors = make([]BuiltinFeverIF, 255)
