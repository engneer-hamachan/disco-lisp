package fever

import (
	"disco/base"
	"disco/printer"
	"fmt"
)

type Princ struct{}

func NewPrinc() BuiltinFeverIF {
	return &Princ{}
}

func init() {
	BuiltinExecutors[base.PRINC] = NewPrinc()
}

func (p *Princ) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	s := VM.PeekStack()

	printer.Print(s, true)
	fmt.Println("")

	return pc, nil
}
