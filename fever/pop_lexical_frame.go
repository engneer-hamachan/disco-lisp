package fever

import (
	"disco/base"
)

type PopLexicalFrame struct{}

func NewPopLexicalFrame() BuiltinFeverIF {
	return &PopLexicalFrame{}
}

func init() {
	BuiltinExecutors[base.POP_LEXICAL_FRAME] = NewPopLexicalFrame()
}

func (p *PopLexicalFrame) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	env.PopStack()

	return pc, nil
}
