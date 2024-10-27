package fever

import (
	"disco/base"
	"disco/printer"
	"fmt"
)

type Print struct{}

func NewPrint() BuiltinFeverIF {
	return &Print{}
}

func init() {
	BuiltinExecutors[base.PRINT] = NewPrint()
}

func (p *Print) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	s := VM.PeekStack()

	printer.Print(s, false)
	fmt.Println("")

	return pc, nil
}
