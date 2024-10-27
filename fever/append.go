package fever

import (
	"disco/base"
	"disco/predicater"
	"fmt"
)

type Append struct{}

func NewAppend() BuiltinFeverIF {
	return &Append{}
}

func init() {
	BuiltinExecutors[base.APPEND] = NewAppend()
}

func (a *Append) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	cadr := VM.PopStack()
	car := VM.PopStack()

	if car.Val != "list-data" && car.Type != base.NIL {
		return pc, fmt.Errorf(
			"%v is invalid type %v want list.",
			car.Val,
			base.TypeToString(car.Type),
		)
	}

	s := a.Append(car, cadr)

	VM.PushStack(s)

	return pc, nil
}

func (a *Append) Append(car *base.S, cdr *base.S) *base.S {
	if predicater.Nilp(car) {
		return cdr
	}

	return base.Cons(car.GetCar(), a.Append(car.GetCdr(), cdr))
}
