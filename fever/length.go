package fever

import (
	"disco/base"
)

type Length struct{}

func NewLength() BuiltinFeverIF {
	return &Length{}
}

func init() {
	BuiltinExecutors[base.LENGTH] = NewLength()
}

func (g *Length) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	s := VM.PopStack()

	switch s.Type {
	case base.LIST:
		if s.GetCar().Type == base.NIL && s.GetCdar() == nil {

			VM.PushStack(base.MakeInt(int64(0)))
			return pc, nil
		}

		ct := 1
		cdr := s.GetCdr()

		for {
			if cdr.Type == base.NIL {
				VM.PushStack(base.MakeInt(int64(ct)))
				return pc, nil
			}

			ct += 1

			cdr = cdr.GetCdr()
		}

	case base.STRING:
		VM.PushStack(base.MakeInt(int64(len(s.Val.(string)))))
		return pc, nil

	default:
		VM.PushStack(base.MakeInt(int64(0)))
		return pc, nil
	}
}
