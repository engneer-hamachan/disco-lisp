package fever

import (
	"disco/base"
)

type PushLexicalFrame struct{}

func NewPushLexicalFrame() BuiltinFeverIF {
	return &PushLexicalFrame{}
}

func init() {
	BuiltinExecutors[base.PUSH_LEXICAL_FRAME] = NewPushLexicalFrame()
}

func (p *PushLexicalFrame) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	env.PushStack()

	return pc, nil
}
