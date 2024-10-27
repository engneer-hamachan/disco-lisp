package fever

import (
	"disco/base"
)

type Car struct{}

func NewCar() BuiltinFeverIF {
	return &Car{}
}

func init() {
	BuiltinExecutors[base.CAR] = NewCar()
}

func (c *Car) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	s := VM.PopStack()

	VM.PushStack(s.GetCar())

	return pc, nil
}
