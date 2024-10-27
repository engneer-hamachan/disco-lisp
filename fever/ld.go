package fever

import (
	"disco/base"
	"fmt"
)

type Ld struct{}

func NewLd() BuiltinFeverIF {
	return &Ld{}
}

func init() {
	BuiltinExecutors[base.LD] = NewLd()
}

func (l *Ld) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	for {
		pc += 1
		s := codes[pc].(*base.S)

		switch s.Type {
		case base.SYMBOL:
			v, ok := env.GetSymbolValue(s)
			if ok {
				if v.Type == base.FUNC {
					f := DefineFunctions[v.Val.(string)]
					VM.PushStack(f)

					break
				}

				VM.PushStack(v)

				break
			}

			return pc, fmt.Errorf("unbound variable %s", s.Val.(string))

		default:
			VM.PushStack(s)
		}

		if len(codes) > pc+1 && codes[pc+1].(int) == base.LD {
			pc += 1
			continue
		}

		break
	}

	return pc, nil
}
